package utils

import (
	"strings"
	"time"
)

// TrimSpace 去除字符串首尾空格
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// GetCurrentTime 获取当前时间
func GetCurrentTime() time.Time {
	return time.Now()
}
