package global

import (
	"hexasoftware/lib/prettylog"
	"log"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(prettylog.New())
}
