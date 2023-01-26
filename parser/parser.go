package parser

import (
	"golang.org/x/net/html"
)

type Parser interface {
	GetLinksFromHtmlNode(node *html.Node) []string
	GetTextFromHtmlNode(node *html.Node) []string
}

// GetLinksFromHtmlNode is a function that takes an *html.Node as input and returns a slice of strings representing the URLs of all the links on the webpage.
func GetLinksFromHtmlNode(node *html.Node) []string {
	links := []string{}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val[1:])
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(node)
	return links
}

// GetLinksFromHtmlNode is a function that takes an *html.Node as input and returns a slice of strings representing the text on the webpage.
func GetTextFromHtmlNode(node *html.Node) []string {
	paragraphs := []string{}
	var searchForParagraphs func(*html.Node)
	searchForParagraphs = func(n *html.Node) {
		if n.Type == html.TextNode && n.Parent.Data == "p" {
			paragraphs = append(paragraphs, n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchForParagraphs(c)
		}
	}

	searchForParagraphs(node)
	return paragraphs
}
