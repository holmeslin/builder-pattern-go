// ctorstat 掃描一個 Go 專案,列出三件事,給簡報的「痛點」頁用真實數字:
//
//  1. 參數超過 -min 個的 New* 函式
//  2. 名字以 Builder 結尾的型別
//  3. 頂層的 With* 函式(Functional Options 的痕跡)
//
// 用法:go run ./cmd/ctorstat -min 5 /path/to/your/repo
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type report struct {
	bigCtors []string
	builders []string
	withs    []string
}

func scan(root string, minParams int) (report, error) {
	var r report
	fset := token.NewFileSet()
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if path != root && (name == "vendor" || name == "testdata" || strings.HasPrefix(name, ".")) {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		f, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
		if err != nil {
			return nil // 解析不了的檔案略過
		}
		rel, _ := filepath.Rel(root, path)
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				name := d.Name.Name
				n := countParams(d.Type.Params)
				pos := fmt.Sprintf("%s:%d", rel, fset.Position(d.Pos()).Line)
				if strings.HasPrefix(name, "New") && n > minParams {
					r.bigCtors = append(r.bigCtors, fmt.Sprintf("%-40s %2d 個參數  %s", name, n, pos))
				}
				if d.Recv == nil && strings.HasPrefix(name, "With") {
					r.withs = append(r.withs, fmt.Sprintf("%-40s %s", name, pos))
				}
			case *ast.GenDecl:
				for _, s := range d.Specs {
					if ts, ok := s.(*ast.TypeSpec); ok && strings.HasSuffix(ts.Name.Name, "Builder") {
						r.builders = append(r.builders, fmt.Sprintf("%-40s %s:%d", ts.Name.Name, rel, fset.Position(ts.Pos()).Line))
					}
				}
			}
		}
		return nil
	})
	sort.Strings(r.bigCtors)
	sort.Strings(r.builders)
	sort.Strings(r.withs)
	return r, err
}

// countParams 把 func(a, b string, c int) 算成 3 個。
func countParams(fl *ast.FieldList) int {
	n := 0
	for _, f := range fl.List {
		if len(f.Names) == 0 {
			n++
		} else {
			n += len(f.Names)
		}
	}
	return n
}

func (r report) print(w io.Writer, minParams int) {
	section := func(title string, items []string) {
		fmt.Fprintf(w, "## %s:%d 個\n", title, len(items))
		for _, it := range items {
			fmt.Fprintln(w, "  "+it)
		}
	}
	section(fmt.Sprintf("參數超過 %d 個的 New* 函式", minParams), r.bigCtors)
	section("*Builder 型別", r.builders)
	section("頂層 With* 函式", r.withs)
}

func main() {
	minParams := flag.Int("min", 5, "參數超過幾個算「很多」")
	flag.Parse()
	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}
	r, err := scan(root, *minParams)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	r.print(os.Stdout, *minParams)
}
