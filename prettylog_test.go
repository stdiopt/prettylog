package prettylog_test

import (
	"testing"

	"github.com/gohxs/prettylog"
)

// Dumb test
func TestPrettylog(t *testing.T) {
	log := prettylog.New("Test")

	log.Println("ok!")
}
