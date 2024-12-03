package tailf

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"loagent/kafka"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

type TailTask struct {
	path       string
	topic      string
	instance   *tail.Tail
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init()
	return tailObj
}

func (t *TailTask) init() {
	config := tail.Config{
		Poll:      true,
		Follow:    true,
		ReOpen:    true,
		MustExist: false,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
	}

	go t.run()
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task: %s_%s done", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			//kafka.SendToKafka(t.topic, line.Text)
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}

//---------------------------------------------------

func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}

func ReadLog() {
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tailObj.Lines
		if !ok {
			fmt.Println("tail file close reopen, filename:%s\n", tailObj.Filename)
			continue
		}
		fmt.Println("msg:", line.Text)
	}
}

func Init0(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}

	tailObj, err = tail.TailFile(filename, config)

	if err != nil {
		panic(err)
	}
	return
}
