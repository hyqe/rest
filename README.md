# rest


```bash
go get github.com/hyqe/rest
```


```go
type Message struct {
	ID   string
	Body string
}

rest.Start(":8080",
	rest.NewRoute(
		func(r rest.Request) rest.Response {
			return rest.JSON(&Message{
				ID:   r.Vars()["id"],
				Body: "Hello",
			})
		},
		rest.WithMethod("GET"),
		rest.WithPath("/messages/{id}"),
	),
)
```