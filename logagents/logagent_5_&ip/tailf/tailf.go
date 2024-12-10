package tailf

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"loagent/kafka"
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
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
