package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// parseLine 解析一行数据并返回字段切片
func parseLine(line string) []string {
	return strings.Fields(line)
}

// parseAddressPort 解析地址和端口
func parseAddressPort(addressPort string) (string, int, error) {
	parts := strings.Split(addressPort, ":")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid address:port format")
	}

	ipHex := parts[0]
	portHex := parts[1]

	ip := parseHexIP(ipHex)
	port, err := strconv.ParseInt(portHex, 16, 32)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port: %v", err)
	}

	return ip, int(port), nil
}

// parseHexIP 将十六进制IP地址转换为人类可读的形式
func parseHexIP(hexIP string) string {
	ipBytes := make([]byte, 4)
	for i := 0; i < 4; i++ {
		byteHex := hexIP[i*2 : i*2+2]
		byteValue, _ := strconv.ParseUint(byteHex, 16, 8)
		ipBytes[3-i] = byte(byteValue) // 反转字节顺序
	}
	return net.IP(ipBytes).String()
}

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // 跳过标题行

	for scanner.Scan() {
		line := scanner.Text()
		fields := parseLine(line)
		if len(fields) < 2 {
			continue
		}

		localAddress, localPort, err := parseAddressPort(fields[1])
		if err != nil {
			fmt.Println("Error parsing local address:", err)
			continue
		}

		remoteAddress, remotePort, err := parseAddressPort(fields[2])
		if err != nil {
			fmt.Println("Error parsing remote address:", err)
			continue
		}

		fmt.Printf("Local: %s:%d, Remote: %s:%d\n", localAddress, localPort, remoteAddress, remotePort)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
