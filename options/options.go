// Package options 是 Functional Options 的對照組(Rob Pike 2014、Dave Cheney 2014)。
// 必填放一般參數,選填每個一個 Option;Option 回傳 error,所以也能驗證。
package options

import (
	"fmt"
	"time"
)

type Server struct {
	addr    string
	port    int
	timeout time.Duration
}

// Option 修改 Server;回傳 error 讓 New 能拒絕不合法的值。
type Option func(*Server) error

func WithPort(p int) Option {
	return func(s *Server) error {
		if p <= 0 || p > 65535 {
			return fmt.Errorf("port %d 無效", p)
		}
		s.port = p
		return nil
	}
}

func WithTimeout(d time.Duration) Option {
	return func(s *Server) error {
		s.timeout = d
		return nil
	}
}

// New:預設值集中在這裡,不給 option 也能用。
func New(addr string, opts ...Option) (*Server, error) {
	s := &Server{addr: addr, port: 8080, timeout: 30 * time.Second}
	for _, o := range opts {
		if err := o(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *Server) String() string {
	return fmt.Sprintf("%s:%d (timeout %v)", s.addr, s.port, s.timeout)
}
