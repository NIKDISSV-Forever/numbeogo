package web

import (
	"golang.org/x/net/html"
	"strings"
)

func GetText(node *html.Node) string {
	textParts := make([]string, 0)
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.TextNode {
			textParts = append(textParts, n.Data)
		}
	}
	return strings.Join(textParts, "")
}
func GetAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
