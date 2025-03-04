# message-push

## 技术栈
| 内容       | 实现                                          |
|----------|---------------------------------------------|
| swap消息推送 |https://bark.day.app/#/?id=bark<br/> |
| unstake推送 | https://bark.day.app/#/?id=bark   |

## 功能介绍

### 1. Swap消息推送
该功能使用Bark实现推送，主要用于监控和推送swap交易的相关消息。通过配置和运行相应的服务，可以实时获取和推送swap交易信息。

### 2. Unstake推送
该功能使用Bark实现推送，主要用于监控和推送unstake操作的相关消息。通过配置和运行相应的服务，可以实时获取和推送unstake操作信息。

## 配置文件说明

### app_config.json
该文件用于配置项目的相关参数，具体配置项说明如下：

- `barkAPIURLs`: 用于发送通知的API URL列表。
- `lastBlockNumber`: 上次处理的区块号。
- `currentTxHashes`: 当前处理的交易哈希列表。

示例配置：
```json
{
  "barkAPIURLs": [
    "https://api.day.app/iuizSoSLLvtMTZhhmuWetY"
  ],
  "lastBlockNumber": "21969425",
  "currentTxHashes": [
    "0xb1fc1bef4e15d94ba06e0e16d1df79b7b7ae09d1b820a0219c933c1d094a320a",
    "0x351d20960377717fe7503367baba41b0a4f793cfa9177980207a7166c3c1eabf"
  ],
  "limitPrice": 1
}
```
## 本地开发启动
```shell
- go run main.go 启动 服务
```
## windows下编译
- 打开powershell ,cd 到项目目录
- 执行以下命令
```shell
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o message-push main.go
```
## 上传到服务
- 编译完成后，会在项目目录下生成一个message-push文件与app_config.json，将该文件上传到服务器65.109.68.159根目录下
- 上传至服务器后设置权限
```shell
chmod +x message-push
chmod +x app_config.json

```
## 启动服务
```shell
nohup ./message-push --config app_config.json &
```
## 查看日志,根目录下执行
```shell
cd logs
tail -100f message_push_output.log
```

