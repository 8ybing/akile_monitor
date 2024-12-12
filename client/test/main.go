package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// 配置结构体
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func tcping(host string, port string) (int64, error) {
	// 检查host是否为空
	if host == "" {
		return 0, fmt.Errorf("主机地址不能为空")
	}

	// 检查port是否为空
	if port == "" {
		return 0, fmt.Errorf("端口号不能为空")
	}

	// 解析端口号
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return 0, fmt.Errorf("端口号转换错误: %v", err)
	}

	// 检查端口范围
	if portNum < 1 || portNum > 65535 {
		return 0, fmt.Errorf("端口号 %d 超出有效范围", portNum)
	}

	// 组合完整的地址
	address := net.JoinHostPort(host, port)

	// 记录开始时间
	startTime := time.Now()

	// 增加详细的连接诊断
	fmt.Printf("正在连接: %s\n", address)

	// 设置连接超时时间为10秒
	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		return 0, fmt.Errorf("连接错误详情: %v", err)
	}
	defer conn.Close()

	// 计算延迟时间（毫秒）
	latency := time.Since(startTime).Milliseconds()

	return latency, nil
}

// 读取配置文件
func readConfig(filename string) (*Config, error) {
	// 解析文件路径，处理特殊字符和相对路径
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("解析文件路径错误: %v", err)
	}

	// 打开文件
	file, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("打开文件错误: %v", err)
	}
	defer file.Close()

	// 解析JSON
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("解析JSON错误: %v", err)
	}

	return &config, nil
}

func main() {
	// 读取配置文件 - 可以是相对或绝对路径
	configPath := `D:\code\Go\akile_monitor\client\test\config.json`

	config, err := readConfig(configPath)
	if err != nil {
		fmt.Printf("读取配置文件错误: %v\n", err)
		os.Exit(1)
	}

	// 执行TCPing
	latency, err := tcping(config.Host, config.Port)
	if err != nil {
		fmt.Printf("TCPing失败: %v\n", err)
		os.Exit(1)
	}

	// 输出延迟
	fmt.Printf("TCPing %s:%s 延迟: %d ms\n", config.Host, config.Port, latency)
}
