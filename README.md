prettylog
============

Simple go package to make log.Logger prettier:

```bash
$ go get github.com/gohxs/prettylog
```

```go
import (
	"log"
	"github.com/gohxs/prettylog"
)

func main() {
	prettylog.Global()
	log.Println("Hello")
}
```
