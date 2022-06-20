package utils

import (
	"bytes"
	"fmt"
	"runtime"
)

//PanicTrace 系统级异常调用栈追踪
func PanicTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\r\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s,%d(0x%x)\r\n", file, line, pc)
	}
	return buf.String()
}
