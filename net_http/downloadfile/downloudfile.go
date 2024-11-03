package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadfile(url string, filename string) {
	r, _ := http.Get(url)
	defer func() { _ = r.Body.Close() }()

	space, _ := os.Create(filename)
	defer func() { _ = space.Close() }()
	n, _ := io.Copy(space, r.Body)
	fmt.Println(n)
}

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	fmt.Printf(" \r进度： %.2f%%", float32(r.Current*10000/r.Total)/100)

	return
}

func downloadProcess(url, filename string) {
	r, _ := http.Get(url)
	defer func() { _ = r.Body.Close() }()
	space, _ := os.Create(filename)
	defer func() { _ = space.Close() }()

	reader := &Reader{
		Reader: r.Body,
		Total:  r.ContentLength,
	}
	_, _ = io.Copy(space, reader)

}

func main() {
	url := "https://th.bing.com/th?id=OIP.GsHAeJ5pOxB6vQ1OVoC7LgHaHa&w=80&h=80&c=1&vt=10&bgcl=3f5cca&r=0&o=6&dpr=1.5&pid=5.1"
	filename := "acg.png2"
	downloadProcess(url, filename)
}
