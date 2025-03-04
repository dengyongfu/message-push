package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// FetchAndParse 发送HTTP GET请求并将响应解析到目标结构体中
func FetchAndParse(url string, target interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var resp *http.Response
	var err error

	for i := 0; i < 3; i++ {
		resp, err = client.Get(url)
		if err == nil {
			break
		}
		fmt.Printf("尝试 %d 失败: %v. 重试中...\n", i+1, err)
		time.Sleep(time.Second)
	}

	if resp == nil {
		return fmt.Errorf("响应为空")
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("非预期的HTTP状态码 %d: %s", resp.StatusCode, string(body))
	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败: %v", err)
	}

	// 解析JSON到目标结构体
	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("解析JSON失败: %v", err)
	}

	return nil
}
