package webscraper

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type WebScraper interface {
	GetHTMLNodeFromUrl(url string) (*html.Node, error)
}

func GetHTMLNodeFromUrl(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return nil, fmt.Errorf("resp status non 200, status: %s ", resp.Status)
	}

	return html.Parse(resp.Body)
}
