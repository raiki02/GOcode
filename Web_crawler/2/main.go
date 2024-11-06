package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func main() {

	Turl := "https://cn.bing.com/"
	resp, err := http.Get(Turl)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("err", resp.Status, resp.StatusCode)
	}
	defer resp.Body.Close()

	// fmt.Println("resp = ", resp)
	// fmt.Println("resp.ContentLength", resp.ContentLength)
	// println("resp.Header", resp.Header)
	// println("resp.Proto", resp.Proto)
	// println("resp.Request", resp.Request)
	// println("resp.TLS", resp.TLS)
	// println("resp.Trailer", resp.Trailer)
	//fmt.Println("resp.Body = ", resp.Body)

	docs, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println("docs = ", docs)

	baseURL, err := url.Parse(Turl)
	if err != nil {
		panic(err)

	}
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link, err := baseURL.Parse(attr.Val)
					if err == nil {
						links = append(links, link.String())
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(docs)

	for v, link := range links {
		fmt.Println(v, link)
	}
}
