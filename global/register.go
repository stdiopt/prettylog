package global

import (
	"hxs/prettylog"
	"log"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(prettylog.New())
}
