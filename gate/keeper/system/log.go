package system

import (
	"log/syslog"
)

var sysLog *syslog.Writer

func InitLoger() {
	// 加载log输出 todo

}

func Loger() *syslog.Writer {
	return sysLog
}
