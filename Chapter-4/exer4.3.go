package main

import "fmt"

func reverse(ints [6]int){
	for i := 0; i < len(ints)/2; i++{
		end := len(ints)- i - 1
		ints[i], ints[end] = ints[end], ints[i]
	}
}

func main(){
	ints := [6]int{0,1,2,3,4,5}
	reverse(ints)
	fmt.Println("%d", ints)
}