package main

import (
	"fmt"
	"strings"
)

// CompareTwoStrings reports whether two strings are anagrams of each other
func CompareTwoStrings(s1, s2 string) bool{
	var yes bool
	if len(s1) == len(s2) {
		for i := range s2{
			if strings.Contains(s1, string(s2[i])) {
				yes = true
			} else {
				yes = false
				break
			}
		}
		return yes
	}
	return false
}

func main(){
	fmt.Println(CompareTwoStrings("care", "d"))
}