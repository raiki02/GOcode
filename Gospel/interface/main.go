package main

import (
	"fmt"
)

func test() {
	// fmt.Sprintln("Hello, World!")
	// a := os.Stdout //type a = *os.File
	// b := io.Writer(a)
	// fmt.Fprintln(b, "Hello, World!")
	// b  &{0xc0000a4248}
	// &b  0xc0000280a0
	// fmt.Println(&b)
	// c := abc{i: 1}
	// fmt.Fprintln(c, "Hello, World!")
	// fmt.Println(c, &c) //{} &{} ??
}

type abc struct {
	i int
}

func (a abc) Write(p []byte) (n int, err error) {
	return 0, nil
}

type byteCounter int

func (c *byteCounter) Write(p []byte) (n int, err error) {
	*c += byteCounter(len(p))
	return len(p), nil
}

func main() {
	var c byteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, \n %s ", name)
}
