package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// IPInfo 定义 JSON 结构体
type IPInfo struct {
	Addr string `json:"addr"`
}

var (
	IPZonePool = make(map[string]string)
	IPZoneLock sync.Mutex
)

// GetIPLocation 获取 IP 地址的地理位置信息
func GetIPLocation(ip string) (string, error) {
	// 是否存在历史IP
	IPZoneLock.Lock()
	CurrentIPZone, ok := IPZonePool[ip]
	IPZoneLock.Unlock()
	if ok {
		return CurrentIPZone, nil
	}

	client := &http.Client{Timeout: time.Second * 5}
	newRequest, err := http.NewRequest("GET", fmt.Sprintf("https://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true", ip), nil)
	// 发送 HTTP GET 请求
	resp, err := client.Do(newRequest)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	} else {
		defer resp.Body.Close()
	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 将 GBK 编码转换为 UTF-8
	utf8Body, err := GbkToUtf8(body)
	if err != nil {
		return "", fmt.Errorf("编码转换失败: %v", err)
	}

	// 解析 JSON 数据
	var ipInfo IPInfo
	err = json.Unmarshal(utf8Body, &ipInfo)
	if err != nil {
		return "", fmt.Errorf("解析 JSON 失败: %v", err)
	} else {
		ipInfo.Addr = strings.Trim(ipInfo.Addr, " ")
		if len(ipInfo.Addr) > 0 {
			// 更新IP信息
			IPZoneLock.Lock()
			IPZonePool[ip] = ipInfo.Addr
			IPZoneLock.Unlock()
		}
	}
	// 返回 addr 字段
	return ipInfo.Addr, nil
}

// GbkToUtf8 将 GBK 编码的字节切片转换为 UTF-8 编码的字节切片
func GbkToUtf8(gbk []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(gbk), simplifiedchinese.GBK.NewDecoder())
	utf8, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return utf8, nil
}
