package popcount

import (
	"testing"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//PopCountLoop Using a for loop
func PopCountLoop(x uint64) int {
	num := 0
	for i := 0; i < 8; i++ {
		num += int(pc[byte(x>>(i*8))])
	}
	return num
}

//PopCountByShifting popcount by shifting
func PopCountByShifting(x uint64) int {
	num := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			num++
		}
	}
	return num
}

//PopCountByClearing popcount by clearing the right most non-zero bit of x
func PopCountByClearing(x uint64) int {
	num := 0
	for x != 0 {
		x = x & (x - 1)
		num++
	}
	return num
}

//BenchmarkPopCount benchmark the PopCount Func
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

//BenchmarkPopCountLoop benchmark the PopCountLoop Func
func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1234567890ABCDEF)
	}
}

//BenchmarkPopCountByShifting benchmark the PopCountByShifting Func
func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

//BenchmarkPopCountByClearing benchmark the PopCountByClearing Func
func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}
