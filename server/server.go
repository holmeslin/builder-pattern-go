// Package server 是情境 2「服務設定」:Go 最常見的 Fluent Builder。
// 方法回傳 *Builder 才能串下去;錯誤先收著,Build() 一次回報。
package server

import (
	"errors"
	"fmt"
	"time"
)

// Server 是成品。
type Server struct {
	Port              int
	Timeout           time.Duration
	CertFile, KeyFile string
}

// Builder 保存半成品與途中發現的錯誤。
type Builder struct {
	s    Server
	errs []error
}

// NewBuilder 回傳一個帶預設值的 Builder。
func NewBuilder() *Builder {
	return &Builder{s: Server{Port: 8080, Timeout: 30 * time.Second}}
}

func (b *Builder) Port(p int) *Builder {
	if p <= 0 || p > 65535 {
		b.errs = append(b.errs, fmt.Errorf("port %d 無效", p))
	}
	b.s.Port = p
	return b // 回傳自己,才能一直點下去
}

func (b *Builder) Timeout(d time.Duration) *Builder {
	if d <= 0 {
		b.errs = append(b.errs, fmt.Errorf("timeout %v 無效", d))
	}
	b.s.Timeout = d
	return b
}

func (b *Builder) TLS(cert, key string) *Builder {
	b.s.CertFile, b.s.KeyFile = cert, key
	return b
}

// Build 是那道門:單一欄位的錯、跨欄位的錯,都在這裡一次回報。
func (b *Builder) Build() (*Server, error) {
	err := errors.Join(b.errs...)
	if (b.s.CertFile == "") != (b.s.KeyFile == "") {
		err = errors.Join(err, errors.New("cert 和 key 要一起設定"))
	}
	if err != nil {
		return nil, err
	}
	s := b.s // 複製一份:之後再改 builder,不會動到已經建好的 Server
	return &s, nil
}
