package compare

import (
	"fmt"
	"time"

	"builderdemo/options"
	"builderdemo/server"
)

func Example_threeWays() {
	// 1. struct literal:直接填,零成本,但沒有驗證
	srv1 := server.Server{Port: 9090, Timeout: 5 * time.Second}
	bad := server.Server{Port: 0} // 編譯器不會叫,執行時才出事
	fmt.Println(srv1.Port, bad.Port)

	// 2. Functional Options:預設值集中,Option 可以回傳 error
	srv2, err := options.New("0.0.0.0",
		options.WithPort(9090), options.WithTimeout(5*time.Second))
	fmt.Println(srv2, err)

	// 3. Builder:一步一步組,Build() 一次驗證
	srv3, err := server.NewBuilder().
		Port(9090).Timeout(5 * time.Second).Build()
	fmt.Println(srv3.Port, err)
	// Output:
	// 9090 0
	// 0.0.0.0:9090 (timeout 5s) <nil>
	// 9090 <nil>
}
