// Package global helper package
// a package intended to be used as
// import _ "github.com/gohxs/prettylog/global"
package global

import (
	"github.com/gohxs/prettylog"
)

func init() {
	prettylog.Global()
}
