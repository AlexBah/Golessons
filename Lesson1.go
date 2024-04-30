package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	stringInFile := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countsLines(os.Stdin, counts, stringInFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countsLines(f, counts, stringInFile)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s\n", stringInFile[line])
			//			for line2, line3 := range stringInFile {
			//				if line == line3 {
			//					fmt.Printf("%s\t%s\n", line, line2)
			//				}
			//			}
			fmt.Printf("----------------\n")
		}
	}
}

func countsLines(f *os.File, counts map[string]int, stringInFile map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		stringInFile[line] = stringInFile[line] + f.Name() + " "
	}
}
