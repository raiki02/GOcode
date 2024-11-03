package main

import (
	"net/http"

	"github.com/PuerKitBio/goquery"
)

func main() {
	spider()
}

func spider() {
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		"https://movie.douban.com/top250",
		nil,
	)
	if err != nil {
		panic(err)
	req.Header.Set("referer", "https://www.bing.com/")  
	
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() { _ = resp.Body.Close() }()

}


v := make(map[string]string)