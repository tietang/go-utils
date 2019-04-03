package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"path"
)

//logrus 代码日志文件，函数和代码行位置输出hook
//用法：
//lfh := NewLogLineNumHook()
//lfh.EnableFileLogLine = true
//lfh.EnableLogFuncName = true
//log.AddHook(lfh)
type LineNumLogrusHook struct {
	//启用文件名称log
	EnableFileNameLog bool
	//启用函数名称log
	EnableFuncNameLog bool
}

func NewLineNumLogrusHook() *LineNumLogrusHook {
	return &LineNumLogrusHook{
		EnableFileNameLog: true,
		EnableFuncNameLog: true,
	}
}

func (hooks LineNumLogrusHook) Levels() []log.Level {
	return log.AllLevels
}

func (hook *LineNumLogrusHook) Fire(entry *log.Entry) error {

	if entry.HasCaller() {
		var (
			file, function string
			line           int
		)
		frame := entry.Caller
		line = frame.Line
		function = frame.Function
		dir, filename := path.Split(frame.File)
		f := path.Base(dir)
		file = fmt.Sprintf("%s/%s", f, filename)
		if hook.EnableFileNameLog && hook.EnableFuncNameLog {
			entry.Message = fmt.Sprintf("[%s(%s:%d)] %s", function, file, line, entry.Message)
		}
		//router/route_table.go(43)
		if hook.EnableFileNameLog && !hook.EnableFuncNameLog {
			entry.Message = fmt.Sprintf("[%s(%d)] %s", file, line, entry.Message)
		}
		//microservice-gateway/v1/router.(*RouteTable).AddRoutePattern(43)
		if !hook.EnableFileNameLog && hook.EnableFuncNameLog {
			entry.Message = fmt.Sprintf("[%s(%d)] %s", function, line, entry.Message)
		}
	}

	return nil
}
