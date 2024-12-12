package main

import (
	"log"
	"net"
	"strings"
)

func getip() string {
	// 存储IP地址的切片
	var ipAddresses []string

	// 获取本地网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error getting interfaces:", err)
		return ""
	}

	for _, iface := range interfaces {
		// 只处理已经启用且没有被回环的接口
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				log.Println("Error getting addresses:", err)
				continue
			}
			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok {
					// 检查是否为IPv4地址
					if ipNet.IP.To4() != nil {
						ipAddresses = append(ipAddresses, ipNet.IP.String())
					}
				}
			}
		}
	}

	// 使用逗号拼接IP地址
	return strings.Join(ipAddresses, ",")
}
