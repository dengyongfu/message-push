package logic

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"log"
	"log/slog"
	"math/big"
	"messag-push/currentcap"
	"net/http"
	"net/url"
	"time"
)

// var networkURL = "https://mainnet.infura.io/v3/80fadbe87b54446f94fbade514eec050"
var networkURL = "https://eth-mainnet.g.alchemy.com/v2/_NivBkHC4lNF3jh-d9mxTPPwKjpScV4K"
var contractAddrStr = "0xAA732c9c110A84d090a72da230eAe1E779f89246"
var tokenAddrStr = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599" // 需要传入的 token 地址
var unstakeAlertConfigUrl = "https://yapi.s.capital/mock/66/unibtc/unstake/alert"

// Notification 是通知服务的接口
type Notification interface {
	Send(message string) error
}

// UnstakeContractLogic 封装监控逻辑
type UnstakeContractLogic struct {
	client       *ethclient.Client
	contract     *currentcap.Currentcap
	contractAddr common.Address
}

// 3. 获取配置

var LastCap = big.NewInt(0)

func init() {
	// 初始化配置
	err := FetchUnstakeAlertConfig()
	if err != nil {
		log.Fatalf("Failed to fetch unstake alert config: %v", err)
	}

}

func CurrentCapTask() error {
	// 1.获取合约
	logic, err := NewUnstakeContractLogic(networkURL, contractAddrStr)
	if err != nil {
		log.Fatalf("Failed to initialize UnstakeContractLogic: %v", err)
		return nil
	}

	// 2.调用合约方法获取cap
	cap, err := logic.GetCurrentCap(common.HexToAddress(tokenAddrStr))
	if err != nil {
		log.Fatalf("Failed to get current cap: %v", err)
		return nil
	}

	slog.Info("cap", "LastCap init", LastCap.String())
	//4. 对配置进行比较判断,是否发送通知
	if !IsUnstakeAlertConfigEmpty(unstakeAlertConfig) && unstakeAlertConfig.Open {
		flag := logic.checkCurrentCap(cap, unstakeAlertConfig.CapThreshold)
		if flag {
			// 5.发送通知
			message := FormatCurrentMessage(cap, &unstakeAlertConfig.CapThreshold)
			title := "WBTC Cap 提醒"
			// 对 title 和 message 进行 URL 编码
			encodedTitle := url.QueryEscape(title)
			encodedMessage := url.QueryEscape(message)

			for _, baseURL := range getBarkAPIURLs() {
				// 拼接 URL
				fullURL := fmt.Sprintf("%s/%s/%s?call=1&level=critical", baseURL, encodedTitle, encodedMessage)
				slog.Info("Notification sent test", "fullURL", fullURL)
				resp, err := http.Get(fullURL)
				if err != nil || resp == nil {
					slog.Error("Failed to send notification to device", "url", fullURL, "error", err)
					continue
				}
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					slog.Error("Notification failed", "url", fullURL, "status", resp.Status)
				} else {
					slog.Info("Notification sent successfully", "url", fullURL)
				}
			}
		}
	}
	// 最后再赋值一次
	LastCap = cap
	slog.Info("cap", "LastCap end", LastCap.String())
	return nil
}

func IsUnstakeAlertConfigEmpty(config UnstakeAlertConfig) bool {
	slog.Info("IsUnstakeAlertConfigEmpty", "UnstakeAlertConfig", config)
	return !config.Open && config.CapThreshold == 0.0
}

func FormatCurrentMessage(currentCap *big.Int, capThreshold *float64) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	readableTime := time.Now().In(loc).Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("%s 当前可用余额: %.8f >= %.8f", readableTime, ConvertToWBTC(currentCap), *capThreshold)
	return message
}

// NewUnstakeContractLogic 初始化并返回 UnstakeContractLogic 实例
func NewUnstakeContractLogic(networkURL string, contractAddrStr string) (*UnstakeContractLogic, error) {
	client, err := ethclient.Dial(networkURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial network")
	}
	contractAddr := common.HexToAddress(contractAddrStr)
	contract, err := currentcap.NewCurrentcap(contractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind contract")
	}

	return &UnstakeContractLogic{
		client:       client,
		contract:     contract,
		contractAddr: contractAddr,
	}, nil
}

// GetCurrentCap 获取当前解约上限
func (l *UnstakeContractLogic) GetCurrentCap(token common.Address) (*big.Int, error) {
	return l.contract.GetCurrentCap(&bind.CallOpts{}, token)
}

// CheckCurrentCap 执行容量检查
func (l *UnstakeContractLogic) checkCurrentCap(currentCap *big.Int, capThreshold float64) bool {
	slog.Info("交易量信息：", "当前容量", ConvertToWBTC(currentCap), "阈值容量", capThreshold, "上次容量", ConvertToWBTC(LastCap))
	// 比较
	if currentCap.Cmp(ConvertFloatToBigInt(capThreshold)) > 0 && ConvertFloatToBigInt(capThreshold).Cmp(LastCap) > 0 {
		slog.Info("警告！", "当前交易量：", ConvertToWBTC(currentCap), "阈值", capThreshold)
		return true
	}
	return false
}
func ConvertToWBTC(currentCap *big.Int) float64 {
	// 将 currentCap 转换为 WBTC 浮点数（假设单位为 1e-8）
	wbtcFloat, _ := new(big.Float).Quo(
		new(big.Float).SetInt(currentCap), // 将 big.Int 转换为 big.Float
		big.NewFloat(1e8),                 // 除以 1e8
	).Float64() // 转换为 float64

	return wbtcFloat
}

// ConvertFloatToBigInt 将 capThreshold 转换为 big.Int
func ConvertFloatToBigInt(capThreshold float64) *big.Int {
	// 将 capThreshold 转换为 uint256（乘以 1e8）
	scale := big.NewFloat(1e8) // 1e8
	capThresholdBig := new(big.Float).Mul(big.NewFloat(capThreshold), scale)

	// 转换为 big.Int
	uint256 := new(big.Int)
	capThresholdBig.Int(uint256)

	return uint256
}
