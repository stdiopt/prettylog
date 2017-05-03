package prettylog

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

type PrettyLogWritter struct {
	prefix   string
	lastTime time.Time
	counter  int64
}

func NewWriter(prefix string) *PrettyLogWritter {
	return &PrettyLogWritter{prefix, time.Now(), 0}
}

func (p *PrettyLogWritter) Write(b []byte) (int, error) {

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

	timeDiff := time.Since(p.lastTime)

	var fduration float64 = float64(timeDiff.Nanoseconds()) / 1000000.0

	msg := fmt.Sprintf("[%d:\033[34m%s\033[0m (\033[33m%s:%d\033[0m) <\033[35m%s\033[0m> \033[90m+%.2f/ms\033[0m]: %s",
		p.counter,
		time.Now().Format("2006-01-02 15:04:05"),
		fname,
		line,
		p.prefix,
		fduration,
		string(b),
	)
	p.lastTime = time.Now()
	p.counter++

	return fmt.Print(msg)
}

func New(prefix string) *log.Logger {
	return log.New(NewWriter(prefix), "", 0)
}

func Global() {
	log.SetFlags(0)
	log.SetOutput(NewWriter(""))
}
