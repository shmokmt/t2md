package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	parsedURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Failed to parse URL:", err)
		os.Exit(1)
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		fmt.Println("Failed to get URL:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println("Failed to create document from response body:", err)
		os.Exit(1)
	}

	title := doc.Find("title").Text()
	fmt.Printf("[%s](%s)\n", title, parsedURL.String())
}
