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
	"time"
)

const graphAPIURL = "https://api.studio.thegraph.com/query/100116/contract_3e2f0/version/latest"
const storageFile = "storage.json"

// 定义多个 Bark API 地址
var barkAPIURLs = []string{
	//"https://api.day.app/kVXVy7PwcQvo2pWTs4QPTQ/",
	"https://api.day.app/iuizSoSLLvtMTZhhmuWetY/%E4%BA%A4%E6%98%93%E6%8F%90%E9%86%92/",
	//"https://api.day.app/UjHSr5Mn2aUpjCee6b2Nkg/%E4%BA%A4%E6%98%93%E6%8F%90%E9%86%92/",
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

type GraphResponse struct {
	Data struct {
		Swaps []Swap `json:"swaps"`
	} `json:"data"`
}

type Storage struct {
	LastBlockNumber string   `json:"lastBlockNumber"`
	CurrentTxHashes []string `json:"currentTxHashes"`
}

var storageData Storage

func init() {
	loadStorage()
}

func loadStorage() {
	file, err := os.Open(storageFile)
	if err != nil {
		// 初始化默认存储数据
		slog.Error("Error opening storage file", "error", err)
		storageData = Storage{LastBlockNumber: "21612681", CurrentTxHashes: []string{"0xccce6256453e517062bb4cfb74494a0bdb2fefa793f75d3d31cf041d76bf99fd"}}
		saveStorage()
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&storageData)
	if err != nil {
		slog.Error("Error decoding storage data", "error", err)
		storageData = Storage{LastBlockNumber: "0", CurrentTxHashes: []string{}}
	}
}

func saveStorage() {
	file, err := os.Create(storageFile)
	if err != nil {
		slog.Error("Error creating storage file", "error", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&storageData)
	if err != nil {
		slog.Error("Error encoding storage data", "error", err)
	}
}

// GraphTask 执行任务
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

	slog.Info("GraphTask", "swaps", swaps)

	// 处理新数据
	var newTxHashes []string
	for _, swap := range swaps {
		if !contains(storageData.CurrentTxHashes, swap.TransactionHash) {
			err = sendNotification(swap)
			if err != nil {
				slog.Error("Error sending notification", "error", err)
			} else {
				newTxHashes = append(newTxHashes, swap.TransactionHash)
			}
		}
	}

	// 更新存储
	if len(swaps) > 0 {
		storageData.LastBlockNumber = swaps[0].BlockNumber
		storageData.CurrentTxHashes = newTxHashes
		saveStorage()
	}
	return nil
}

func fetchSwaps() ([]Swap, error) {
	pageSize := 50
	startBlock := parseLastBlockNumber()
	var allSwaps []Swap

	for {
		// 格式化 GraphQL 查询字符串
		query := fmt.Sprintf(queryTemplate, pageSize, startBlock)

		// 创建请求体
		requestBody, err := json.Marshal(map[string]string{
			"query": query,
		})
		if err != nil {
			slog.Error("Failed to create request body", "error", err)
			return nil, err
		}

		// 创建 HTTP 请求
		req, err := http.NewRequest("POST", graphAPIURL, bytes.NewBuffer(requestBody))
		if err != nil {
			slog.Error("Failed to create HTTP request", "error", err)
			return nil, err
		}

		// 设置请求头
		req.Header.Set("Content-Type", "application/json")

		// 执行请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			slog.Error("Failed to execute request", "resp", resp, "error", err)
			return nil, err
		}
		defer resp.Body.Close()

		// 读取响应体
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			slog.Error("Failed to read response body", "error", err)
			return nil, err
		}

		// 解析响应
		var graphResponse GraphResponse
		err = json.Unmarshal(body, &graphResponse)
		if err != nil {
			slog.Error("Failed to parse response body", "error", err)
			return nil, err
		}

		slog.Info("fetchSwaps", "graphResponse", graphResponse)
		// 如果没有新数据，退出循环
		if len(graphResponse.Data.Swaps) == 0 {
			break
		}

		// 追加当前查询的数据
		allSwaps = append(allSwaps, graphResponse.Data.Swaps...)

		// 更新起始 BlockNumber 为当前结果中的最大 BlockNumber
		newStartBlock, err := strconv.Atoi(graphResponse.Data.Swaps[len(graphResponse.Data.Swaps)-1].BlockNumber)
		if err != nil {
			slog.Error("Failed to parse BlockNumber", "error", err)
			return nil, err
		}
		startBlock = newStartBlock

		// 如果查询到的数据少于分页大小，说明没有更多数据了
		if len(graphResponse.Data.Swaps) < pageSize {
			break
		}
	}
	return allSwaps, nil
}

func parseLastBlockNumber() int {
	blockNumber, err := strconv.Atoi(storageData.LastBlockNumber)
	if err != nil {
		slog.Error("Failed to parse stored BlockNumber, defaulting to 0", "error", err)
		return 21584546
	}
	return blockNumber
}

func sendNotification(swap Swap) error {
	timestamp, err := strconv.ParseInt(swap.BlockTimestamp, 10, 64)
	if err != nil {
		timestamp = time.Now().Unix()
		slog.Error("Failed to parse blockTimestamp", "error", err)
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")

	readableTime := time.Unix(timestamp, 0).In(loc).Format("2006-01-02 15:04:05")
	slog.Info("New swap detected", "blockNumber",
		swap.BlockNumber, "transactionHash", swap.TransactionHash, "blockTimes", readableTime, "btcPrice", swap.BtcPrice)

	slog.Info("sendNotification", "amount0:", swap.Amount0, "amount1:", swap.Amount1)
	message := FormatSwap(&swap)
	if message == "" {
		return nil
	}

	// 遍历所有 Bark API 地址进行推送
	for _, baseURL := range barkAPIURLs {
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

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// FormatSwap formats the Swap event into the desired string format
func FormatSwap(swap *Swap) string {
	// Convert amounts from strings directly to big.Float
	amount0Float, _ := new(big.Float).SetString(swap.Amount0)
	amount1Float, _ := new(big.Float).SetString(swap.Amount1)

	// Initialize variables for amountIn, amountOut, tokenIn, and tokenOut
	var amountIn, amountOut *big.Float
	var tokenIn, tokenOut string

	// Determine the direction of the swap based on the sign of amount0
	if amount0Float.Sign() < 0 { // Selling token0 (e.g., WBTC) to buy token1 (e.g., UNIBTC)
		amountIn = amount1Float                      // amount0 is negative, so it's the amount we are selling
		amountOut = new(big.Float).Neg(amount0Float) // amount1 is the amount we are buying
		tokenIn = "WBTC"
		tokenOut = "UNIBTC"
	} else { // Selling token1 (e.g., UNIBTC) to buy token0 (e.g., WBTC)
		amountIn = amount0Float                      // amount0 is positive, so it's the amount we are selling
		amountOut = new(big.Float).Neg(amount1Float) // amount1 is the amount we are buying
		tokenIn = "UNIBTC"
		tokenOut = "WBTC"
	}

	// Calculate volume in USD using the WBTC price
	defaultPrice := 100000.0
	wbtcPrice := big.NewFloat(defaultPrice)
	if swap.BtcPrice != "" {
		if parsedPrice, _, err := new(big.Float).Parse(swap.BtcPrice, 10); err == nil {
			wbtcPrice = parsedPrice
		} else {
			slog.Error("Failed to parse btcPrice", "error", err)
		}
	}

	vol := new(big.Float).Mul(amountIn, wbtcPrice)

	// Format the output, ensuring that amounts are rounded to 5 decimal places
	amountInStr := new(big.Float).Quo(amountIn, big.NewFloat(1e8)).Text('f', 5)
	amountOutStr := new(big.Float).Quo(amountOut, big.NewFloat(1e8)).Text('f', 5)

	volStr := new(big.Float).Quo(vol, big.NewFloat(1e8)).Text('f', 2)

	timestamp, err := strconv.ParseInt(swap.BlockTimestamp, 10, 64)
	if err != nil {
		return ""
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	readableTime := time.Unix(timestamp, 0).In(loc).Format("2006-01-02 15:04:05")

	// Ensure the output is accurate and formatted
	return fmt.Sprintf("%s  %s %s -> %s %s Vol: $%s",
		readableTime, amountInStr, tokenIn, amountOutStr, tokenOut, volStr)
}
