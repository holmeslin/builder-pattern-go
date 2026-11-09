package overkill

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"
)

type measure struct {
	defLines  int // 定義的程式碼行數(不含空行、註解、package、import、use* 函式)
	defFuncs  int // 定義的函式/方法數(不含 use* 函式)
	callLines int // use* 函式裡呼叫端那一句佔幾行
}

func measureFile(path string) measure {
	src, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	var m measure
	skip := map[int]bool{}
	for _, d := range f.Decls {
		fd, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if strings.HasPrefix(fd.Name.Name, "use") {
			from, to := fset.Position(fd.Pos()).Line, fset.Position(fd.End()).Line
			for l := from; l <= to; l++ {
				skip[l] = true
			}
			stmt := fd.Body.List[0]
			m.callLines = fset.Position(stmt.End()).Line - fset.Position(stmt.Pos()).Line + 1
			continue
		}
		m.defFuncs++
	}
	for i, l := range strings.Split(string(src), "\n") {
		l = strings.TrimSpace(l)
		if skip[i+1] || l == "" || strings.HasPrefix(l, "//") ||
			strings.HasPrefix(l, "package ") || strings.HasPrefix(l, "import ") {
			continue
		}
		m.defLines++
	}
	return m
}

// 投影片「反例」頁的數字就是這段 Output,由 go test 逐字比對。
func Example_cost() {
	lit := measureFile("literal.go")
	bld := measureFile("builder.go")
	fmt.Printf("struct literal:定義 %d 行、%d 個函式,呼叫端 %d 行\n",
		lit.defLines, lit.defFuncs, lit.callLines)
	fmt.Printf("硬套 Builder:多 %d 行、多 %d 個函式,呼叫端 %d 行\n",
		bld.defLines, bld.defFuncs, bld.callLines)
	// Output:
	// struct literal:定義 5 行、0 個函式,呼叫端 1 行
	// 硬套 Builder:多 17 行、多 5 個函式,呼叫端 5 行
}

// 兩種寫法產出的 Config 完全一樣:Builder 沒有多擋下任何錯誤。
func TestSameResult(t *testing.T) {
	if useLiteral() != useBuilder() {
		t.Fatal("兩種寫法結果不同")
	}
}
