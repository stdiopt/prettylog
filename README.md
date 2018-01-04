prettylog
============

Simple go package to make log.Logger prettier:

```bash
$ go get github.com/gohxs/prettylog
```

```go
import (
	"github.com/gohxs/prettylog"
	"log"
)

func main() {
	prettylog.Global()
	log.Println("Hello")
}
```
