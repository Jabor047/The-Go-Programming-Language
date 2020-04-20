
// There is No sha384 instead the programs prints sha256 and sha512
// and prints an error if no width / length is specified.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var width = flag.Int("w", 256, "hash width (256 or 512)")

func main(){
	flag.Parse()
	var function func(b []byte) []byte
	switch *width{
	case 256:
		function = func(b []byte) []byte{
			h := sha256.Sum256(b)
			return h[:]
		}
	case 512:
		function = func(b []byte)[]byte{
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("Unexpected width specified")
	}
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%x\n", function(b))
}