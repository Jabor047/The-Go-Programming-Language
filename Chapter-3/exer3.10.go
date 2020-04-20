package main

import (
	"fmt"
	"bytes"
)


func comma(s string)string{
	
	var buf bytes.Buffer
	
	for i, v := range s{
		if i != 0 && i % 3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}
	return buf.String()
}

func main(){
	fmt.Println(comma("negative"))
}