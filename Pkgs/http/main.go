package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func _canonicalheaderkey() {
	str := "accept-encoding"
	str2 := http.CanonicalHeaderKey(str)
	fmt.Println(str2) //"Accept-Encoding"
	str = "aaa"
	str2 = http.CanonicalHeaderKey(str)
	fmt.Println(str2) //Aaa
	str = "a-a-a"
	str2 = http.CanonicalHeaderKey(str)
	fmt.Println(str2) //A-A-A
}

func _detectcontenttype() {
	str := "asdfghjkl"
	str2 := http.DetectContentType([]byte(str))
	fmt.Println(str2) //text/plain; charset=utf-8
	str = "id:01"
	str2 = http.DetectContentType([]byte(str))
	fmt.Println(str2) //text/plain; charset=utf-8
	res, _ := http.Get("http://httpbin.org/get")
	body, _ := io.ReadAll(res.Body)
	fmt.Println(http.DetectContentType(body)) //text/plain; charset=utf-8
}

func _parsehttpcersion() {
	mj, mn, ok := http.ParseHTTPVersion("HTTP/1.0")
	fmt.Println(mj, mn, ok) //1 0 true
	mj, mn, ok = http.ParseHTTPVersion("HTP/1.2")
	fmt.Println(mj, mn, ok) //0 0 false
}

func _parsetime() {
	t, err := http.ParseTime("12:00:09")
	fmt.Print(t, err) //0001-01-01 00:00:00 +0000 UTC parsing time "12:00:09" as "Mon Jan _2 15:04:05 2006": cannot parse "12:00:09" as "Mon"2024-11
	res, _ := http.Get("http://httpbin.org/get")
	date := res.Header.Get("Date")
	t, err = http.ParseTime(date)
	fmt.Print(t, err) //-10 16:12:21 +0000 UTC <nil>
	t, err = http.ParseTime("Mon Jan 2 15:04:05 2006")
	fmt.Print(t, err) //-10 16:24:40 +0000 UTC <nil>   2006-01-02 15:04:05 +0000 UTC <nil>
	t, err = time.Parse("11:11:12", "Mon Jan 2 15:04:05 2006")
	fmt.Println(t, err)
	t, err = http.ParseTime("1")
}

func makegetmethod() {
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080",
		nil,
	)
	client := http.Client{}
	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Printf("%s", body)

}
func makepostmethod() {
	str := strings.NewReader("post body")
	req, _ := http.NewRequest(
		http.MethodPost,
		"http://localhost:8080",
		str,
	)
	client := http.Client{}
	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Printf("%s", body)

}

func openserver() {
	http.ListenAndServe(":8080", nil)
}

func handle() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)
		w.Header().Set("Content-Type", "test/plain")
		w.Header().Set("X-Custom-Header", "Nice-Cum-from-1145141919810")
		switch r.Method {
		case http.MethodGet:
			io.WriteString(w, "teg")
		case http.MethodPost:
			io.WriteString(w, "psot")
			body, _ := io.ReadAll(r.Body)
			io.WriteString(w, fmt.Sprintf("body =  %s", string(body)))
		}

	}
	http.HandleFunc("/", handler)
}

func usemostfunc() {

}

func handle_test() {

}

func main() {
	fmt.Println("helo world")
}

// var mu sync.Mutex
// var count int

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/count", counter)
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }

// // handler echoes the Path component of the requested URL.
// func handler(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

// // counter echoes the number of calls so far.
// func counter(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	fmt.Fprintf(w, "Count %d\n", count)
// 	mu.Unlock()
// }
