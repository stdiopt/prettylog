package prettylog_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/gohxs/prettylog"
)

// Dumb test
func TestPrettylog(t *testing.T) {

	log := prettylog.New("Test")

	log.Println("ok!")
	log.Println("ok!")

	log2 := prettylog.New("")
	log2.Println("Test")
	log.Println("ok!")
}

func TestMultipleWriters(t *testing.T) {
	origTime := time.Date(1980, time.Month(10), 12, 10, 0, 0, 0, time.Local)
	p := monkey.Patch(time.Now, func() time.Time { return origTime })
	defer p.Unpatch()

	out1 := bytes.NewBuffer([]byte{})
	out2 := bytes.NewBuffer([]byte{})

	log := prettylog.New("test", out1, out2, os.Stderr)
	log.Println("Hello test")

	if out1.Len() == 0 {
		t.Fatal("Nothing was recorded")
	}

	if out1.String() != out2.String() {
		t.Fatal("Log does not match")
	}
}
