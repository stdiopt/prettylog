package prettylog

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

type PrettyLogWritter struct {
	lastTime time.Time
	counter  int64
}

func New() *PrettyLogWritter {
	return &PrettyLogWritter{time.Now(), 0}
}

func (this *PrettyLogWritter) Write(b []byte) (int, error) {

	/*{
		for i := 0; i < 6; i++ {
			ptr, _, _, _ := runtime.Caller(i)
			fname := runtime.FuncForPC(ptr).Name()
			fmt.Println("Stack:", fname)
		}
	}*/

	ptr, _, line, _ := runtime.Caller(3)
	tname := runtime.FuncForPC(ptr).Name()
	li := strings.LastIndex(tname, "/")
	fname := tname[li+1:]

	timeDiff := time.Since(this.lastTime)

	var fduration float64 = float64(timeDiff.Nanoseconds()) / 1000000.0

	msg := fmt.Sprintf("[%d:\033[34m%s\033[0m (\033[33m%s:%d\033[0m) \033[90m+%.2f/ms\033[0m]: %s",
		this.counter,
		time.Now().Format("2006-01-02 15:04:05"),
		fname,
		line,
		fduration,
		string(b),
	)
	this.lastTime = time.Now()
	this.counter++

	return fmt.Print(msg)
}

func CreateLogger() *log.Logger {
	return log.New(New(), "", 0)
}
