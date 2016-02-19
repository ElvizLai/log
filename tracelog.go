package log

import (
	"fmt"
	"time"
)

// TraceIn and TraceOut use in function in and out,reduce code line
// Example:
//	func test() {
//		user := User{Name: "zhangsan", Age: 21, School: "xayddx"}
//		service := "verification.GetVerifiCode"
//		defer log.TraceOut(log.TraceIn("12345", service, "user:%v", user))
//		....
//	}

func TraceIn(traceID string, service string, format string, v ...interface{}) (string, string, time.Time) {
	startTime := time.Now()
	Std.Output(traceID, Linfo, 2, fmt.Sprintf(service+", "+format, v...))
	return traceID, service, startTime
}

func TraceOut(traceID string, service string, startTime time.Time) {
	Std.Output(traceID, Linfo, 2, fmt.Sprintf("finished "+service+", took%v", time.Since(startTime)))
}
