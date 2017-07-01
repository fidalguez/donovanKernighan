package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, fileMap := range counts {
		//fmt.Printf("%s\n",line)
		var N int
		var audit string
		for fileName, n := range fileMap {
			N += n
			audit += fileName + " (x" + strconv.Itoa(n) + ")\t"
		}
		if N > 1 {
			fmt.Printf("%d\t%s\t", N, line)
			fmt.Printf("%s\n",audit)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			subMap := make(map[string]int)
			counts[input.Text()] = subMap
		}
		counts[input.Text()][f.Name()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}