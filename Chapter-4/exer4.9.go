package main

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	freq := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan(){
		word := input.Text()
		freq[word]++
	}
	if input.Err() != nil{
		fmt.Fprintln(os.Stderr, input.Err())
		os.Exit(1)
	}
	for word, count := range freq{
		fmt.Printf("%-30s %d\n", word, count)
	}
}