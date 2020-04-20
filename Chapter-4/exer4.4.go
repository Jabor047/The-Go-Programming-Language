package main

import "fmt"

func rotateInts(ints []int){
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints) - 1] = first
}

func main(){
	ints := []int{2,3,4,5,6,7}
	rotateInts(ints)
	fmt.Println(ints)
}