package logic

import (
	"log"
	"log/slog"
	"messag-push/utils"
	"sync"
)

// UnstakeAlertConfig 结构体
type UnstakeAlertConfig struct {
	Open         bool    `json:"open"`
	CapThreshold float64 `json:"capThreshold"`
}

// SwapAlertConfig 结构体
type SwapAlertConfig struct {
	Open            bool `json:"open"`
	VolumeThreshold int  `json:"volumeThreshold"` // 限制 BTC 价格
}

var (
	unstakeAlertConfig      UnstakeAlertConfig
	unstakeAlertConfigMutex sync.RWMutex // 配置读写锁
	swapAlertConfig         SwapAlertConfig
	swapAlertConfigMutex    sync.RWMutex // 配置读写锁
)

func init() {
	err := FetchAlertConfig()
	if err != nil {
		slog.Error("FetchAlertConfig", "Failed to initialize FetchAlertConfig", err)
	}

}

// FetchUnstakeAlertConfig .获取配置
func FetchUnstakeAlertConfig() error {
	var config UnstakeAlertConfig
	err := utils.FetchAndParse(unstakeAlertConfigUrl, &config)
	if err != nil {
		log.Fatalf("UnstakeAlertConfig获取和解析数据失败: %v", err)
		return err
	}
	unstakeAlertConfigMutex.Lock()

	unstakeAlertConfig = config
	defer unstakeAlertConfigMutex.Unlock()
	return nil
}

// FetchSwapAlertConfig .获取配置
func FetchSwapAlertConfig() error {
	var sAlertConfig SwapAlertConfig
	err := utils.FetchAndParse(swapAlertConfigUrl, &sAlertConfig)
	if err != nil {
		log.Fatalf("SwapAlertConfig获取和解析数据失败: %v", err)
		return err
	} else {
		swapAlertConfigMutex.Lock()
		swapAlertConfig = sAlertConfig
		defer swapAlertConfigMutex.Unlock()
		return nil
	}
}

func FetchAlertConfig() error {
	err := FetchUnstakeAlertConfig()
	if err != nil {
		log.Fatalf("FetchUnstakeAlertConfig 获取和解析数据失败: %v", err)
	}
	err = FetchSwapAlertConfig()
	if err != nil {
		log.Fatalf("FetchSwapAlertConfig 获取和解析数据失败: %v", err)
	}
	return nil
}
