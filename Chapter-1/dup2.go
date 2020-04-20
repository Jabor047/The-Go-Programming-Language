package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filenames := make([]string, len(files))
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for fil := range filenames {
		fmt.Printf("%s", fil)
	}
}

func countLines(f *os.File, counts map[string]int, filenames []string) {
	arr := []string{f.Name()}
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for _, n := range counts {
		if n > 1 {
			filenames = append(filenames, arr...)
		}
		break
	}
}
