package title1

import (
	"fmt"
	"strings"
	"net/http"
	"golang.org/x/net/html"
	"../links"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	
	// Check Content-Type is HTML (e.g "text/html; charset=utf-8")
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;"){
		resp.Body.Close()
		return fmt.Errorf("%s hat type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node){
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	links.ForEachNode(doc, visitNode, nil)
	return nil
}

func titleWithDiffer(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;"){
		return fmt.Errorf("%s hat type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node){
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	links.ForEachNode(doc, visitNode, nil)
	return nil
}