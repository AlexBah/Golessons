// display arguments of command line
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	for i, arg := range os.Args[0:] {
		s = strconv.Itoa(i) + " " + arg
		fmt.Println(s)
	}
	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())
}
