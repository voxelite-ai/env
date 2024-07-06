# Env
 
Type safe env variables

```sh
get get -u github.com/voxelite-ai/env
```


## Usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/voxelite-ai/env"
)

func main() {
	port := env.Int64("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	addr := fmt.Sprintf(":%d", port)

	http.ListenAndServe(addr, mux)
}
```
