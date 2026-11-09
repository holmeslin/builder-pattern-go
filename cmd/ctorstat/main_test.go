package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestScanSample(t *testing.T) {
	r, err := scan("testdata/sample", 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.bigCtors) != 1 || !strings.Contains(r.bigCtors[0], "NewReport") ||
		!strings.Contains(r.bigCtors[0], "10 個參數") {
		t.Fatalf("bigCtors = %v", r.bigCtors)
	}
	if len(r.builders) != 1 || !strings.HasPrefix(r.builders[0], "QueryBuilder") {
		t.Fatalf("builders = %v", r.builders)
	}
	if len(r.withs) != 1 || !strings.HasPrefix(r.withs[0], "WithPort") {
		t.Fatalf("withs = %v", r.withs)
	}
	r.print(os.Stdout, 5)
}

// 簡報 p.2 的程式碼放在 testdata/sample,確認它真的能編譯、go vet 通過。
func TestSampleCompiles(t *testing.T) {
	if testing.Short() {
		t.Skip("需要呼叫 go vet")
	}
	if out, err := exec.Command("go", "vet", "./testdata/sample").CombinedOutput(); err != nil {
		t.Fatalf("go vet 失敗:\n%s", out)
	}
}
