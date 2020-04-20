package main

import "fmt"

func elimDuplicates(strings []string)[]string{
	dup := 0
	for _, s := range strings {
		if strings[dup] == s{
			continue
		}
		dup++
		strings[dup] = s
	}
	return strings[:dup+1]
} 

func main(){
	strs := []string{"Ben","Ben", "Kev", "Sam","Sam", "Ron"}
	fmt.Println("%s", elimDuplicates(strs))
}