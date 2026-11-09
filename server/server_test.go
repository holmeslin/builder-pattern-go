// 用外部測試套件 server_test:程式碼長得跟呼叫端一模一樣(server.NewBuilder())。
package server_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/holmeslin/builder-pattern-go/server"
)

func ExampleBuilder() {
	srv, err := server.NewBuilder().
		Port(9090).
		Timeout(5*time.Second).
		TLS("c.pem", "k.pem").
		Build()
	fmt.Println(srv.Port, srv.Timeout, err)

	_, err = server.NewBuilder().
		Port(0).Timeout(-1).TLS("c.pem", "").
		Build()
	fmt.Println(err)
	// Output:
	// 9090 5s <nil>
	// port 0 無效
	// timeout -1ns 無效
	// cert 和 key 要一起設定
}

// 質疑:「Build() 回傳的 Server,之後改 builder 會不會被影響?」
func TestBuildReturnsCopy(t *testing.T) {
	b := server.NewBuilder().Port(9090)
	s1, _ := b.Build()
	b.Port(8443)
	if s1.Port != 9090 {
		t.Fatalf("已建好的 Server 被改了: Port = %d", s1.Port)
	}
}

// 質疑:「忘了呼叫 Build() 怎麼辦?」
// 答:*Builder 不是 *Server,要用到成品的地方,編譯器會擋。
func TestForgetBuildDoesNotCompile(t *testing.T) {
	if testing.Short() {
		t.Skip("需要呼叫 go build")
	}
	out, err := exec.Command("go", "build", "./testdata/forgetbuild").CombinedOutput()
	if err == nil {
		t.Fatal("預期編譯失敗,卻成功了")
	}
	if !strings.Contains(string(out), "cannot use") {
		t.Fatalf("預期型別錯誤,實際輸出:\n%s", out)
	}
	t.Logf("編譯器輸出:\n%s", out)
}
