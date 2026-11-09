package docgen

import (
	"fmt"
	"strings"
	"testing"
)

// 投影片上的「輸出」就是這段 Example 的 Output,go test 會逐字比對。
func ExampleBuildMonthly() {
	data := Monthly{Month: "10", Summary: "一切正常", Notes: "無"}

	for _, b := range []ReportBuilder{
		&HTMLBuilder{}, &MarkdownBuilder{},
		&StatsBuilder{},
	} {
		BuildMonthly(b, data)
		fmt.Println(b.Result())
	}
	// Output:
	// <h1>10 月報</h1><h2>摘要</h2><p>一切正常</p><h2>備註</h2><p>無</p>
	// # 10 月報
	// ## 摘要
	// 一切正常
	// ## 備註
	// 無
	//
	// 2 個章節
}

func TestHTMLBuilderEscapes(t *testing.T) {
	b := &HTMLBuilder{}
	BuildMonthly(b, Monthly{Month: "10", Summary: "<script>alert(1)</script>"})
	if strings.Contains(b.Result(), "<script>") {
		t.Fatalf("HTML 沒有跳脫: %s", b.Result())
	}
}
