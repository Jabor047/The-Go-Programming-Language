package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"net/http"
	"golang.org/x/net/html"
)

// CountWordsAndImages does a HTTP GET request for the HTML
// document url and returns the number of words and images in it
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("Parsing HTML : %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImages counts the number of images and words in a page
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		words += countWords(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		words, images = countWordsAndImages(c)
	}
	return
}

// countWords counts the number of words in a string 
func countWords(s string) (number int) {
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		number++
	}
	return
}

func main() {
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Words: %d\t Images: %d\n", words, images)
}