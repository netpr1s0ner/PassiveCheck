package runner

import (
	"PassiveCheck/static"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Metadata struct {
	FofaQuery   []string `json:"fofa-query"`
	GoogleQuery []string `json:"google-query"`
	Product     string   `json:"product"`
	ShodanQuery []string `json:"shodan-query"`
	Vendor      string   `json:"vendor"`
	Verified    bool     `json:"verified"`
}

type Info struct {
	Name     string   `json:"name"`
	Author   string   `json:"author"`
	Tags     string   `json:"tags"`
	Severity string   `json:"severity"`
	Metadata Metadata `json:"metadata"`
}

type Matcher struct {
	Type  string   `json:"type"`
	Regex []string `json:"regex"`
	Words []string `json:"words"`
}

type Http struct {
	Method   string    `json:"method"`
	Path     []string  `json:"path"`
	Matchers []Matcher `json:"matchers"`
}

type Item struct {
	ID   string `json:"id"`
	Info Info   `json:"info"`
	Http []Http `json:"http"`
}

var (
	items               []Item
	fingersPrintHub     = make(map[int]*finger)
	fingersPrintHubPath = []string{""}
	FingersFile         = "FingerprintHub.txt"
)

func Run() {
	// 打开文件
	fileContent, err := os.ReadFile(FingersFile)
	if err != nil {
		fileContent = static.FingerPrintV4
		return
	}

	// 解析 JSON 数据
	err = json.Unmarshal(fileContent, &items)
	if err != nil {
		return
	}

	// 打印解析后的数据
	count := 0
	for _, item := range items {
		for _, tmpHttp := range item.Http {

			// 路由
			for _, path := range tmpHttp.Path {
				tmpPath := strings.ReplaceAll(path, "{{BaseURL}}", "")
				if !contains(fingersPrintHubPath, tmpPath) {
					fingersPrintHubPath = append(fingersPrintHubPath, tmpPath)
				}
			}

			for _, matcher := range tmpHttp.Matchers {
				count++
				if matcher.Type == "favicon" {
					for _, icon_hash := range item.Info.Metadata.FofaQuery {
						if strings.Contains(icon_hash, "icon_hash=") {
							fingersPrintHub[count] = &finger{
								Name:    item.Info.Name,
								Type:    faviconType,
								Keyword: []string{icon_hash},
							}
						}
					}
				} else if matcher.Type == "word" {
					fingersPrintHub[count] = &finger{
						Name:    item.Info.Name,
						Type:    bodyType,
						Keyword: matcher.Words,
					}
				} else if matcher.Type == regexType {
					fingersPrintHub[count] = &finger{
						Name:    item.Info.Name,
						Type:    regexType,
						Keyword: matcher.Regex,
					}
				}

			}
		}
	}
}

func fingerprintIdentificationFingerprintHub(r *Runner, body, data []byte, URL, LastURL string) []string {
	var (
		result      []string
		matchType   = bodyType // 匹配类型
		matchString string     // 匹配关键字
	)

	// 计算 favicon
	if strings.Contains(URL, ".ico") || strings.Contains(LastURL, ".ico") || strings.Contains(URL, ".png") || strings.Contains(LastURL, ".png") {
		MMH3Hash, _, err := r.calculateFaviconHashWithRaw(data)
		if err != nil {
			return result
		} else {
			matchString = MMH3Hash
			matchType = faviconType
		}
	} else {
		matchString = string(body)
	}

	for key, value := range fingersPrintHub {
		if value.Type != matchType {
			continue
		}

		// 匹配favicon
		if matchType == faviconType {
			if strings.Contains(value.Keyword[0], matchString) {
				result = append(result, fmt.Sprintf("%d.%s", key, value.Name))
			}
		} else { // 匹配body
			hitFlag := true
			if value.Type == bodyType {
				for _, cKeyWord := range value.Keyword {
					if !strings.Contains(cKeyWord, matchString) {
						hitFlag = false
						break
					}
				}
				if hitFlag {
					result = append(result, fmt.Sprintf("%d.%s", key, value.Name))
				}
			} else if value.Type == regexType { // 正则匹配
				re := regexp.MustCompile(value.Keyword[0])
				if re.FindStringIndex(matchString) != nil {
					result = append(result, fmt.Sprintf("%d.%s", key, value.Name))
				}
			}
		}
	}
	return result
}

func ShowFingerprintHub(searchID int) {
	count := 0
	// 指纹信息 - 1
	for _, item := range items {
		for _, tmpHttp := range item.Http {
			for _, _ = range tmpHttp.Matchers {
				count++
				if count == searchID {
					fmt.Println("指纹信息:")
					fmt.Println(item)
					fmt.Println()
					break
				}
			}
		}
	}

	// 指纹信息-2
	for key, value := range fingers {
		if key == searchID {
			fmt.Println("指纹信息 ", "系统:", value.Name, "匹配类型:", value.Type, "keyword", value.Keyword)
			break
		}
	}
}

func UpdateFingerprintHub(filename string) error {
	resp, err := http.Get("https://github.com/0x727/FingerprintHub/raw/refs/heads/main/service_fingerprint_v4.json")
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		// 将内容写入本地文件
		err = os.WriteFile(filename, content, 0644)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
