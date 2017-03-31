godaemon
========

Run golang app as background program

## Get it：

```
go get github.com/qodrorid/godaemon
```

## Example:

```go
package main

import (
	_ "github.com/icattlecoder/godaemon"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello, golang!\n"))
	})
	log.Fatalln(http.ListenAndServe(":3030", mux))
}
```

## Run it

```
./example -d=true
~$ curl http://127.0.0.1:3030/index
hello, golang!
```

# Enjoy :)