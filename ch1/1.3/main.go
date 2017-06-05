package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""

	start := time.Now()

	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println("elapsed:", time.Since(start))
	fmt.Println(s)

	start = time.Now()
	s = strings.Join(os.Args, sep)
	fmt.Println("elapsed:", time.Since(start))
	fmt.Println(s)
}
