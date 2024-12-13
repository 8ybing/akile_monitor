package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

// TCPPing 执行TCP连接ping并返回延迟
func TCPPing(address string, timeout time.Duration) (time.Duration, error) {
	// 记录开始时间
	start := time.Now()

	// 建立TCP连接
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return 0, fmt.Errorf("连接失败: %v", err)
	}
	defer conn.Close()

	// 计算延迟
	elapsed := time.Since(start)

	return elapsed, nil
}

// ReadConfigAndPing 从配置文件读取地址并执行ping
func ReadConfigAndPing(configPath string, timeout time.Duration) (time.Duration, error) {
	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Errorf("读取配置文件错误: %v", err)
	}

	// 创建一个map来存储配置
	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Errorf("解析JSON错误: %v", err)
	}

	// 从配置中提取tcping_address
	tcpingAddress, ok := config["tcping_address"].(string)
	if !ok {
		fmt.Errorf("无效的tcping_address")
	}

	// 执行TCP Ping
	delay, err := TCPPing(tcpingAddress, timeout)
	if err != nil {
		fmt.Errorf("TCP Ping失败: %v", err)
	}

	// 打印结果
	//fmt.Printf("目标地址: %s\n", tcpingAddress)
	//fmt.Printf("连接延迟: %v\n", delay)

	return delay, nil
}

func main() {
	// 设置超时时间（例如5秒）
	timeout := 5 * time.Second

	// 执行ping
	latency, err := ReadConfigAndPing("D:\\code\\Go\\akile_monitor\\config.json", timeout)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(latency)
}
