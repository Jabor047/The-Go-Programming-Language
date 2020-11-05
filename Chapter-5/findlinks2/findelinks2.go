package main

import (
	"fmt"
	"os"
	"net/http"
	"../html"
	"../findlinks1"
)

func main(){
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		} 
	}
}

// findLinks performs a HTTP GET request for an URL,parses the response
// as HTML and extracts and returns the links
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s, url, resp.status")
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}