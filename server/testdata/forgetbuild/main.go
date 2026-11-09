// 這個程式故意編譯不過:忘了呼叫 Build(),把 *Builder 當成 *Server 用。
// server_test.go 會確認編譯器真的擋下它。
package main

import "github.com/holmeslin/builder-pattern-go/server"

func run(s *server.Server) {}

func main() {
	b := server.NewBuilder().Port(9090)
	run(b) // 忘了 .Build()
}
