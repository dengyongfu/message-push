package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	graphAPIURL = "https://api.studio.thegraph.com/query/100116/contract_3e2f0/version/latest"
	configFile  = "app_config.json" // 合并后的配置文件
)

// 配置文件结构
type Config struct {
	BarkAPIURLs     []string `json:"barkAPIURLs"`     // Bark API 地址列表
	LastBlockNumber string   `json:"lastBlockNumber"` // 上次处理的区块号
	CurrentTxHashes []string `json:"currentTxHashes"` // 当前已处理的交易哈希列表
}

var (
	configData  Config       // 全局配置数据
	configMutex sync.RWMutex // 配置读写锁
)

func init() {
	// 初始化时加载配置
	loadConfig()
	// 启动文件监控
	go watchConfig()
}

// 加载配置文件
func loadConfig() {
	file, err := os.Open(configFile)
	if err != nil {
		slog.Error("Error opening config file, using default config", "error", err)
		// 如果配置文件不存在，使用默认配置
		configData = Config{
			BarkAPIURLs: []string{
				"https://api.day.app/iuizSoSLLvtMTZhhmuWetY/%E4%BA%A4%E6%98%93%E6%8F%90%E9%86%92/",
			},
			LastBlockNumber: "21612681",
			CurrentTxHashes: []string{"0xccce6256453e517062bb4cfb74494a0bdb2fefa793f75d3d31cf041d76bf99fd"},
		}
		saveConfig()
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var newConfig Config
	err = decoder.Decode(&newConfig)
	if err != nil {
		slog.Error("Error decoding config data", "error", err)
		return
	}

	// 更新全局配置
	configMutex.Lock()
	configData = newConfig
	configMutex.Unlock()
}

// 保存配置文件
func saveConfig() {
	file, err := os.Create(configFile)
	if err != nil {
		slog.Error("Error creating config file", "error", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 格式化输出
	err = encoder.Encode(&configData)
	if err != nil {
		slog.Error("Error encoding config data", "error", err)
	}
}

// 监控配置文件变化
func watchConfig() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		slog.Error("Failed to create watcher", "error", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(configFile)
	if err != nil {
		slog.Error("Failed to add config file to watcher", "error", err)
		return
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				slog.Info("Config file modified, reloading...")
				loadConfig() // 配置文件修改时重新加载
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			slog.Error("Watcher error", "error", err)
		}
	}
}

// 获取 Bark API 地址列表
func getBarkAPIURLs() []string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return configData.BarkAPIURLs
}

// 获取上次处理的区块号
func getLastBlockNumber() string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return configData.LastBlockNumber
}

// 获取当前已处理的交易哈希列表
func getCurrentTxHashes() []string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return configData.CurrentTxHashes
}

// 更新上次处理的区块号
func setLastBlockNumber(blockNumber string) {
	configMutex.Lock()
	defer configMutex.Unlock()
	configData.LastBlockNumber = blockNumber
}

// 更新当前已处理的交易哈希列表
func setCurrentTxHashes(txHashes []string) {
	configMutex.Lock()
	defer configMutex.Unlock()
	configData.CurrentTxHashes = txHashes
}

// GraphQL 查询模板
const queryTemplate = `
{
  swaps(first: %d, orderBy: blockNumber, orderDirection: desc, where: {blockNumber_gt: %d}) {
    id
    sender
    recipient
    amount0
    amount1
    sqrtPriceX96
    liquidity
    tick
    blockNumber
    blockTimestamp
    transactionHash
    btcPrice
  }
}`

// Swap 数据结构
type Swap struct {
	ID              string `json:"id"`
	Sender          string `json:"sender"`
	Recipient       string `json:"recipient"`
	Amount0         string `json:"amount0"`
	Amount1         string `json:"amount1"`
	SqrtPriceX96    string `json:"sqrtPriceX96"`
	Liquidity       string `json:"liquidity"`
	Tick            int32  `json:"tick"`
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  string `json:"blockTimestamp"`
	TransactionHash string `json:"transactionHash"`
	BtcPrice        string `json:"btcPrice"`
}

// GraphResponse 数据结构
type GraphResponse struct {
	Data struct {
		Swaps []Swap `json:"swaps"`
	} `json:"data"`
}

// 获取最新的 Swap 数据
func fetchSwaps() ([]Swap, error) {
	pageSize := 50
	startBlock, _ := strconv.Atoi(getLastBlockNumber())
	var allSwaps []Swap

	for {
		query := fmt.Sprintf(queryTemplate, pageSize, startBlock)
		requestBody, err := json.Marshal(map[string]string{"query": query})
		if err != nil {
			slog.Error("Failed to create request body", "error", err)
			return nil, err
		}

		req, err := http.NewRequest("POST", graphAPIURL, bytes.NewBuffer(requestBody))
		if err != nil {
			slog.Error("Failed to create HTTP request", "error", err)
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			slog.Error("Failed to execute request", "error", err)
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			slog.Error("Failed to read response body", "error", err)
			return nil, err
		}

		var graphResponse GraphResponse
		err = json.Unmarshal(body, &graphResponse)
		if err != nil {
			slog.Error("Failed to parse response body", "error", err)
			return nil, err
		}

		if len(graphResponse.Data.Swaps) == 0 {
			break
		}

		allSwaps = append(allSwaps, graphResponse.Data.Swaps...)
		newStartBlock, _ := strconv.Atoi(graphResponse.Data.Swaps[len(graphResponse.Data.Swaps)-1].BlockNumber)
		startBlock = newStartBlock

		if len(graphResponse.Data.Swaps) < pageSize {
			break
		}
	}
	return allSwaps, nil
}

// 发送通知
func sendNotification(swap Swap) error {
	timestamp, _ := strconv.ParseInt(swap.BlockTimestamp, 10, 64)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	readableTime := time.Unix(timestamp, 0).In(loc).Format("2006-01-02 15:04:05")
	slog.Info("New swap detected", "blockNumber", swap.BlockNumber, "transactionHash", swap.TransactionHash, "blockTimes", readableTime, "btcPrice", swap.BtcPrice)

	message := FormatSwap(&swap)
	if message == "" {
		return nil
	}

	for _, baseURL := range getBarkAPIURLs() {
		baseURL = baseURL + message + "?call=1"
		resp, err := http.Get(baseURL)
		if err != nil {
			slog.Error("Failed to send notification to device", "url", baseURL, "error", err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			slog.Error("Notification failed", "url", baseURL, "status", resp.Status)
		} else {
			slog.Info("Notification sent successfully", "url", baseURL)
		}
	}
	return nil
}

// 格式化 Swap 数据
func FormatSwap(swap *Swap) string {
	amount0Float, _ := new(big.Float).SetString(swap.Amount0)
	amount1Float, _ := new(big.Float).SetString(swap.Amount1)

	var amountIn, amountOut *big.Float
	var tokenIn, tokenOut string

	if amount0Float.Sign() < 0 {
		amountIn = amount1Float
		amountOut = new(big.Float).Neg(amount0Float)
		tokenIn = "WBTC"
		tokenOut = "UNIBTC"
	} else {
		amountIn = amount0Float
		amountOut = new(big.Float).Neg(amount1Float)
		tokenIn = "UNIBTC"
		tokenOut = "WBTC"
	}

	wbtcPrice := big.NewFloat(100000.0)
	if swap.BtcPrice != "" {
		if parsedPrice, _, err := new(big.Float).Parse(swap.BtcPrice, 10); err == nil {
			wbtcPrice = parsedPrice
		} else {
			slog.Error("Failed to parse btcPrice", "error", err)
		}
	}

	vol := new(big.Float).Mul(amountIn, wbtcPrice)
	amountInStr := new(big.Float).Quo(amountIn, big.NewFloat(1e8)).Text('f', 5)
	amountOutStr := new(big.Float).Quo(amountOut, big.NewFloat(1e8)).Text('f', 5)
	volStr := new(big.Float).Quo(vol, big.NewFloat(1e8)).Text('f', 2)

	timestamp, err := strconv.ParseInt(swap.BlockTimestamp, 10, 64)
	if err != nil {
		return ""
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	readableTime := time.Unix(timestamp, 0).In(loc).Format("2006-01-02 15:04:05")

	return fmt.Sprintf("%s  %s %s -> %s %s Vol: $%s", readableTime,
		amountInStr, tokenIn, amountOutStr, tokenOut, volStr)
}

// 主任务
func GraphTask() error {
	swaps, err := fetchSwaps()
	if err != nil {
		slog.Error("Error fetching swaps", "error", err)
		time.Sleep(3 * time.Second)
		return err
	}
	if len(swaps) == 0 {
		slog.Info("No new swaps found")
		return nil
	}

	var newTxHashes []string
	for _, swap := range swaps {
		if !contains(getCurrentTxHashes(), swap.TransactionHash) {
			err = sendNotification(swap)
			if err != nil {
				slog.Error("Error sending notification", "error", err)
			} else {
				newTxHashes = append(newTxHashes, swap.TransactionHash)
			}
		}
	}

	if len(swaps) > 0 {
		setLastBlockNumber(swaps[0].BlockNumber)
		setCurrentTxHashes(newTxHashes)
		saveConfig()
	}
	return nil
}

// 判断切片是否包含某个元素
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
