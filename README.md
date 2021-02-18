# rest


```bash
go get github.com/hyqe/rest
```


```go
package main

import (
	"github.com/hyqe/rest"
)

type Message struct {
	ID   string
	Body string
}

func main() {
	routes := []rest.Route{
		rest.NewRoute(
			func(r rest.Request) rest.Response {
				msg := &Message{
					Body: "Hi!",
				}
				r.Vars(&msg)
				return rest.JSON(&msg)
			},
			rest.WithMethod("GET"),
			rest.WithPath("/messages/{id}"),
		),
	}

	rest.Start(":8080", routes...)
}
```

```bash
curl "http://localhost:8080/messages/123"
```