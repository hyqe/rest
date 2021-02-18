# rest


```bash
go get github.com/hyqe/rest
```


```go
type Message struct {
	ID   string
	Body string
}

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
```