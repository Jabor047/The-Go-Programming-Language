package main

import "fmt"

func appendInt(x []int, y ...int) []int{
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x){
		// There is room to grow Extend the slice
		z = x[ :zlen]
	} else{
		//If there is no sufficient space. Allocate a new array.
		//Grow by doubling
		zcap := zlen
		if zcap < 2 * len(x){
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func main(){
	var x,y []int
	for i := 0; i < 10; i++{
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}