package tailf

import "github.com/hpcloud/tail"

type Tailf_mgr struct {
	tailers map[string]*TailTask
}

var (
	mgr    *Tailf_mgr
	Config tail.Config
)

func Init() {
	mgr = &Tailf_mgr{
		tailers: make(map[string]*TailTask, 16),
	}
	Config = tail.Config{
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}
}

func NewTailTask(path string, topic string) {

}
