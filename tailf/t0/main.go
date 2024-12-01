package main

import (
	"github.com/hpcloud/tail"
)

func main() {
	readFile := "./logtest.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile(readFile, config)
	if err != nil {
		panic(err)

	}
	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			println("tail file close reopen, filename:%s\n", tails.Filename)
			continue
		}
		println("msg:", line.Text)

	}
}
