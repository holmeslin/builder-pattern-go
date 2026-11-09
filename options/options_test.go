package options

import (
	"fmt"
	"time"
)

func ExampleNew() {
	s, _ := New("0.0.0.0")
	fmt.Println(s)

	s, _ = New("0.0.0.0", WithPort(9090), WithTimeout(5*time.Second))
	fmt.Println(s)

	_, err := New("0.0.0.0", WithPort(0))
	fmt.Println(err)
	// Output:
	// 0.0.0.0:8080 (timeout 30s)
	// 0.0.0.0:9090 (timeout 5s)
	// port 0 無效
}
