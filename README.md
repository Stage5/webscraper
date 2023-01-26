# Web Scraper and Web Parser Library

This library contains two packages for web scraping: webscraper and parser.

## webscraper
This package contains one function:

### GetHTMLNodeFromUrl(url string) (*html.Node, error)
This function takes a string representing a URL and returns the root *html.Node of the corresponding webpage.

## parser
This package contains two functions:

### GetLinksFromHtmlNode(node *html.Node) ([]string)
This function takes an *html.Node and returns a slice of strings representing the URLs of all the links on the webpage.

### GetTextFromHtmlNode(node *html.Node) ([]string)
This function takes an *html.Node and returns a slice of strings containing the text of the webpage.