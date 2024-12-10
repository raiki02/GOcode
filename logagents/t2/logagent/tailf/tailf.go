package tailf

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
)

type TailTask struct {
	Path   string
	Topic  string
	Cancel context.CancelFunc
}

func begin(path string, topic string) {
	tailObj, err := tail.TailFile(path, Config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}

}

func run() {

}
