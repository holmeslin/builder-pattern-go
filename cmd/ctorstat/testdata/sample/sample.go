// 簡報 p.2「痛點」頁的程式碼就是這個檔案;ctorstat 的測試拿它當樣本。
package sample

import "time"

type Report struct{}

type Section struct{}

// 一開始只有三個參數,後來⋯
func NewReport(title, author string,
	date time.Time, sections []Section,
	footer, pageSize string,
	landscape, withTOC, withCover bool,
	watermark *string) *Report {
	return &Report{}
}

func caller(now time.Time, secs []Section) *Report {
	// 呼叫端長這樣
	r := NewReport("月報", "Holme", now, secs,
		"", "A4", false, true, false, nil)
	return r
}

// 2 個參數:不該被抓到。
func NewSmall(a, b int) *Report { return &Report{} }

type QueryBuilder struct{}

type Server struct{ port int }

type Option func(*Server)

func WithPort(p int) Option { return func(s *Server) { s.port = p } }

// 方法叫 With 開頭不算(有 receiver)。
func (s *Server) WithNothing() {}
