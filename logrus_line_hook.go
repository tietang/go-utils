package utils

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "runtime"
    "strings"
)

//logrus 代码日志文件，函数和代码行位置输出hook
type LogLineNumHook struct {
    //启用文件名称log
    EnableFileNameLog bool
    //启用函数名称log
    EnableFuncNameLog bool
    Skip              int
}

func NewLogLineNumHook() *LogLineNumHook {
    return &LogLineNumHook{
        EnableFileNameLog: true,
        EnableFuncNameLog: true,
        Skip:              5,
    }
}

func (hooks LogLineNumHook) Levels() []log.Level {
    return log.AllLevels
}

func (hook *LogLineNumHook) Fire(entry *log.Entry) error {

    file, function, line := hook.findCaller(hook.Skip)

    if hook.EnableFileNameLog && hook.EnableFuncNameLog {
        entry.Message = fmt.Sprintf("[%s(%s:%d)] [%s]", function, file, line, entry.Message)
    }
    //router/route_table.go(43)
    if hook.EnableFileNameLog && !hook.EnableFuncNameLog {
        entry.Message = fmt.Sprintf("[%s(%d)] %s", file, line, entry.Message)
    }
    //microservice-gateway/v1/router.(*RouteTable).AddRoutePattern(43)
    if !hook.EnableFileNameLog && hook.EnableFuncNameLog {
        entry.Message = fmt.Sprintf("[%s(%d)] %s", function, line, entry.Message)
    }

    return nil
}

func (hook *LogLineNumHook) findCaller(skip int) (string, string, int) {
    var (
        pc       uintptr
        file     string
        function string
        line     int
    )
    for i := 0; i < 10; i++ {
        pc, file, line = hook.getCaller(skip + i)
        if !strings.HasPrefix(file, "logrus/") {
            break
        }
    }
    if pc != 0 && hook.EnableFuncNameLog {
        frames := runtime.CallersFrames([]uintptr{pc})
        frame, _ := frames.Next()
        function = frame.Function
    }

    return file, function, line
}

func (hook *LogLineNumHook) getCaller(skip int) (uintptr, string, int) {
    pc, file, line, ok := runtime.Caller(skip)
    if !ok {
        return 0, "", 0
    }

    n := 0
    for i := len(file) - 1; i > 0; i-- {
        if file[i] == '/' {
            n += 1
            if n >= 2 {
                file = file[i+1:]
                break
            }
        }
    }

    return pc, file, line
}
