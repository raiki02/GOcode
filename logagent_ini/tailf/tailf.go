package tailf

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tails *tail.Tail
)

func Init(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}

	tails, err = tail.TailFile(filename, config)

	if err != nil {
		panic(err)
	}
	return
}

func ReadLog() {
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen, filename:%s\n", tails.Filename)
			continue
		}
		fmt.Println("msg:", line.Text)
	}
}

func ReadChan() <-chan *tail.Line {
	return tails.Lines
}
