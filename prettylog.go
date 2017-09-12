/*Package prettylog Low performance but pretty and coherent log writer */
package prettylog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/gohxs/prettylog/style"
)

var (
	// Style Global log style options
	Style = style.NewDefault(
		style.Options{
			"Counter":  {Color: "\033[37m", UseGlobalPad: true},
			"Message":  {Color: "\033[37m"},
			"Prefix":   {Color: "\033[33m", UseGlobalPad: true},
			"Time":     {Color: "\033[34m"},
			"Duration": {Color: "\033[90m"},
			"File":     {Color: "\033[30m"},
		},
	)
)

//Writter writer struct
type Writter struct {
	prefix   string
	lastTime time.Time
	counter  int64
	writers  []io.Writer

	// Flags??
}

//NewWriter creates a new log writer to be used in log.SetOutput
func NewWriter(prefix string, writers ...io.Writer) *Writter {
	return &Writter{prefix, time.Now(), 0, writers}
}

//Write io.Write implementation that parses the output
func (p *Writter) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	originalLen := len(b)

	parts := strings.Split(string(b), "\n")
	if len(parts) > 2 {
		for _, v := range parts {
			p.Write([]byte(v))
		}
		return originalLen, nil
	}
	msg := parts[0]
	//msg := string(b)
	/*{
		for i := 0; i < 6; i++ {
			ptr, _, _, _ := runtime.Caller(i)
			fname := runtime.FuncForPC(ptr).Name()
			fmt.Println("Stack:", fname)
		}
	}*/

	ptr, file, line, _ := runtime.Caller(3)

	tname := runtime.FuncForPC(ptr).Name()
	//fmt.Println("Tname:", tname)

	method := tname[strings.LastIndex(tname, ".")+1:]
	fname := file[strings.LastIndex(file, "/")+1:]

	timeDiff := time.Since(p.lastTime)

	duration := durationStr(timeDiff)

	prefixStr := method
	if p.prefix != "" {
		prefixStr = fmt.Sprintf("%s", p.prefix)
	}
	/*if !terminal.IsTerminal(int(os.Stderr.Fd())) {
		style.Disabled = true
	}*/
	//msg := fmt.Sprintf("[%d:\033[34m%s\033[0m (\033[33m%s:%d\033[0m) %s\033[90m+%.2f/ms\033[0m]: %s",
	str := fmt.Sprintf("[%s:%s %s]: %s %s %s\n",
		Style.Get("Counter", p.counter),
		Style.Get("Time", time.Now().Format("2006-01-02 15:04:05.000")),
		Style.Get("Prefix", prefixStr),
		Style.Get("Message", msg),

		Style.Get("Duration", duration),
		Style.GetX("File", fmt.Sprintf("%s:%d", fname, line)),
	)

	p.lastTime = time.Now()
	p.counter++

	for _, w := range p.writers {
		n, err := w.Write([]byte(str))
		if err != nil {
			return n, err
		}
	}

	return originalLen, nil
}

// New creates a new log.Logger with a prefix
func New(prefix string, writers ...io.Writer) *log.Logger {
	return log.New(NewWriter(prefix, writers...), "", 0)
}

// Dummy a log.Logger with io.Discard writer
func Dummy() *log.Logger {
	return log.New(ioutil.Discard, "", 0)
}

// Global sets the global log with a prettylog writer
func Global() {
	log.SetFlags(0)
	log.SetOutput(NewWriter(""))
}
