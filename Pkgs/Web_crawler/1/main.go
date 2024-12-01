package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

// fetchHTML 发送 HTTP 请求并返回页面的 HTML 内容
func fetchHTML(targetURL string) (*html.Node, error) {
	resp, err := http.Get(targetURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("无法访问 %s, 状态码: %d", targetURL, resp.StatusCode)
	}

	// 解析 HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// extractLinks 提取 HTML 文档中的所有链接
func extractLinks(doc *html.Node, baseURL *url.URL) []string {
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					// 解析相对链接为绝对链接
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
	f(doc)
	return links
}

func main() {
	var targetURL string
	fmt.Print("请输入你想抓取资源的网站：")
	fmt.Scanln(&targetURL)
	// 获取并解析 HTML 内容
	doc, err := fetchHTML(targetURL)
	if err != nil {
		log.Fatalf("获取页面失败: %v", err)
	}

	// 解析目标 URL，用于处理相对链接
	baseURL, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("无法解析 URL: %v", err)
	}

	// 提取所有链接
	links := extractLinks(doc, baseURL)
	fmt.Println("页面中的链接：")
	for v, link := range links {
		fmt.Println(v, link)
	}
}
