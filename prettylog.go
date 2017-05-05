/* Low performance but pretty and coherent log writer */
package prettylog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	style = NewStyle()
)

type PrettyLogWritter struct {
	prefix   string
	lastTime time.Time
	counter  int64

	// Flags??
}

func NewWriter(prefix string) *PrettyLogWritter {
	return &PrettyLogWritter{prefix, time.Now(), 0}
}

func (p *PrettyLogWritter) Write(b []byte) (int, error) {
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

	ptr, _, line, _ := runtime.Caller(3)
	tname := runtime.FuncForPC(ptr).Name()
	li := strings.LastIndex(tname, "/")
	fname := tname[li+1:]

	timeDiff := time.Since(p.lastTime)

	fdurationSuf := "ms"
	fduration := float64(timeDiff.Nanoseconds()) / 1000000.0
	if fduration > 100 {
		fduration /= 1000
		fdurationSuf = "s"
	}

	prefixStr := fmt.Sprintf("%12s", p.prefix)

	if !terminal.IsTerminal(int(os.Stderr.Fd())) {
		style.Disabled = true
	}
	//msg := fmt.Sprintf("[%d:\033[34m%s\033[0m (\033[33m%s:%d\033[0m) %s\033[90m+%.2f/ms\033[0m]: %s",
	str := fmt.Sprintf("[%s %s]: %s %s %s\n",
		style.Get("Time", time.Now().Format("2006-01-02 15:04:05")),
		style.Get("Prefix", prefixStr),
		style.Get("Message", msg),

		style.Get("Duration", fmt.Sprintf("+%.2f/%s", fduration, fdurationSuf)),
		style.GetX("File", fmt.Sprintf("%s:%d", fname, line)),
	)
	p.lastTime = time.Now()
	p.counter++

	n, err := os.Stderr.Write([]byte(str))
	if err != nil {
		return n, err
	}

	return originalLen, nil
}

func New(prefix string) *log.Logger {
	return log.New(NewWriter(prefix), "", 0)
}

func Global() {
	log.SetFlags(0)
	log.SetOutput(NewWriter(""))
}
