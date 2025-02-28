package runner

import (
	"net/url"
	"strings"
)

func getPathFromURL(rawURL string) (string, error) {
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "http://" + rawURL
	}

	// 解析URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	} else if parsedURL.Path == "" || len(parsedURL.Path) == 0 {
		return "/", nil
	} else {
		return parsedURL.Path, nil
	}
}

// 判断 []string 中是否包含某个字符串
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
