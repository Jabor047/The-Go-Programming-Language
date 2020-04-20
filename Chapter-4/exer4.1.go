package main

import (
	"crypto/sha256"
	"fmt"
)

// popcount returns the population count (number of set bits)
func popCount(b byte)int{
	count := 0
	for ; b != 0; count++{
		b &= b - 1
	}
	return count
}

// bitDiff returns the number of different bits of hash1 and hash2
func bitDiff(hsh1, hsh2 []byte)int{
	num := 0
	for i := 0; i < len(hsh1) || i < len(hsh2); i++ {
		switch {
		case i >= len(hsh1):
			num += popCount(hsh2[i])
		case i >= len(hsh2):
			num += popCount(hsh1[i])
		default:
			num += popCount(hsh1[i] ^ hsh2[i])
		}
	}
	return num
}

func main(){
	hash1 := sha256.Sum256([]byte{1})
	hash2 := sha256.Sum256([]byte{2}) 

	fmt.Println("%d", bitDiff(hash1[:], hash2[:]))
}