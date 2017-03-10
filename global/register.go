package global

import (
	"log"

	prettylog "dev.hexasoftware.com/hxs/prettylog.git"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(prettylog.New())
}
