package tailf

import "loagent/etcd"

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry: logEntryConf,
	}

	for _, logEntry := range logEntryConf {
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
