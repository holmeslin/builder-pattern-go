// Package docgen 是情境 1「文件輸出」:GoF 原版的 Builder。
// 一個介面(ReportBuilder)、多個實作、一個 Director(BuildMonthly)。
// 同一套步驟,換一個 builder 就產出不同的成品。
package docgen

import (
	"fmt"
	"html"
	"strings"
)

// ReportBuilder 定義建構一份報表的步驟。
type ReportBuilder interface {
	SetTitle(title string)
	AddSection(heading, body string)
	Result() string
}

// Monthly 是月報的原始資料。
type Monthly struct{ Month, Summary, Notes string }

// BuildMonthly 是 Director:步驟寫死在這裡
func BuildMonthly(b ReportBuilder, m Monthly) {
	b.SetTitle(m.Month + " 月報")
	b.AddSection("摘要", m.Summary)
	b.AddSection("備註", m.Notes)
}

// HTMLBuilder 產出 HTML。內容一律跳脫,避免 XSS。
type HTMLBuilder struct{ sb strings.Builder }

func (b *HTMLBuilder) SetTitle(t string) {
	fmt.Fprintf(&b.sb, "<h1>%s</h1>", html.EscapeString(t))
}

func (b *HTMLBuilder) AddSection(h, body string) {
	fmt.Fprintf(&b.sb, "<h2>%s</h2>", html.EscapeString(h))
	fmt.Fprintf(&b.sb, "<p>%s</p>", html.EscapeString(body))
}

func (b *HTMLBuilder) Result() string { return b.sb.String() }

// MarkdownBuilder 產出 Markdown。
type MarkdownBuilder struct{ sb strings.Builder }

func (b *MarkdownBuilder) SetTitle(t string) { fmt.Fprintf(&b.sb, "# %s\n", t) }

func (b *MarkdownBuilder) AddSection(h, body string) {
	fmt.Fprintf(&b.sb, "## %s\n%s\n", h, body)
}

func (b *MarkdownBuilder) Result() string { return b.sb.String() }

// StatsBuilder 不產生文字,只數章節。同一套步驟,成品型別完全不同。
type StatsBuilder struct{ sections int }

func (b *StatsBuilder) SetTitle(string) {}

func (b *StatsBuilder) AddSection(_, _ string) {
	b.sections++
}

func (b *StatsBuilder) Result() string {
	return fmt.Sprintf("%d 個章節", b.sections)
}
