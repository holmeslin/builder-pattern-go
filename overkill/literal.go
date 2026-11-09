// Package overkill 是反例:3 個欄位、全部選填、沒有驗證,硬套 Builder 的成本。
// cost_test.go 會實際數這兩個檔案的程式碼行數與函式數。
package overkill

import "time"

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func useLiteral() Config {
	return Config{Host: "db", Port: 5432, Timeout: 3 * time.Second}
}
