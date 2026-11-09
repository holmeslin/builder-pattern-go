package overkill

import "time"

type ConfigBuilder struct{ c Config }

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{}
}

func (b *ConfigBuilder) Host(h string) *ConfigBuilder {
	b.c.Host = h
	return b
}

func (b *ConfigBuilder) Port(p int) *ConfigBuilder {
	b.c.Port = p
	return b
}

func (b *ConfigBuilder) Timeout(d time.Duration) *ConfigBuilder {
	b.c.Timeout = d
	return b
}

// Build 沒有任何驗證:這道門什麼都沒擋。
func (b *ConfigBuilder) Build() Config { return b.c }

func useBuilder() Config {
	return NewConfigBuilder().
		Host("db").
		Port(5432).
		Timeout(3 * time.Second).
		Build()
}
