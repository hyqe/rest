package rest_test

import (
	"github.com/hyqe/rest"
)

// Example ...
func Example() {

	type Message struct {
		ID   string
		Body string
	}

	rest.Start(":8080",
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
	)
}
