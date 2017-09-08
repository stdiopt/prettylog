package prettylog

import (
	"fmt"
	"time"
)

func durationStr(dur time.Duration) string {
	fdurationSuf := "ms"
	fduration := float64(dur.Nanoseconds()) / 1000000.0
	if fduration > 100 {
		fduration /= 1000
		fdurationSuf = "s"
	}

	return fmt.Sprintf("+%.2f/%s", fduration, fdurationSuf)
}
