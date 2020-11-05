package main

import (
	"golang.org/x/net/html"
)

// ElementByTagName finds all elements that have the name tag
func ElementByTagName(doc *html.Node, name...string)[]*html.Node {
	var tags []*html.Node
	if doc.Type == html.ElementNode {
		for _, tag := range name {
			if doc.Data == tag {
				tags = append(tags, doc)
			}
		}
	}
	return tags
}