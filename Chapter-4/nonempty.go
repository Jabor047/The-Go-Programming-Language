package main

import "fmt"

// nonempty returns a slice holding only the non empty 	strings
// The Underlying array is modified during the call
func nonempty(strings []string)[]string{
	i := 0
	for _, s := range strings{
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string{
	out := strings[:0] //zero-length slice of the original
	for _, s := range strings{
		if s != ""{
			out = append(out, s)
		}
	}
	return out	
}

func main(){
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}