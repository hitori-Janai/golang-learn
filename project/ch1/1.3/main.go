//测试三个echo的性能 
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	count := 40000

	start := time.Now()
	var s, sep string
	for index := 0; index < count; index++ {
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	s = ""
	for index := 0; index < count; index++ {
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	for index := 0; index < count; index++ {
		strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

//!-
