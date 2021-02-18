package rest

import (
	"fmt"
	"io"
	"net/http"
)

// Start a REST API
func Start(addr string, routes ...Route) error {
	return http.ListenAndServe(addr, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var req Request
		for _, route := range routes {
			if _, ok := route.Match(req); ok {
				resp := route.Handler(req)
				if resp != nil {
					for k, v := range resp.Headers() {
						rw.Header().Set(k, v)
					}
					rw.WriteHeader(resp.Status())
					fmt.Fprint(rw, resp.Body())
				}
				break
			}
		}
	}))
}

// Response ...
type Response interface {
	Status() int
	Body() []byte
	Headers() map[string]string
}

// NewResponse ..
func NewResponse(opts ...ResponseOption) Response {
	return nil
}

// ResponseOption ..
type ResponseOption func(r Response)

// SetBody ..
func SetBody(b io.Reader) ResponseOption {
	return func(r Response) {

	}
}

// WithJSON ..
func WithJSON(v interface{}) ResponseOption {
	return func(r Response) {

	}
}

// JSON ...
func JSON(v interface{}) Response {
	return nil
}

// Request ...
type Request interface {
	Base() *http.Request
	Vars(interface{})
}

// Handler ...
type Handler func(Request) Response

// Route ..
type Route struct {
	Handler
	Matches
}

// NewRoute ..
func NewRoute(h Handler, m ...MatchOption) Route {
	return Route{
		Handler: h,
		Matches: m,
	}
}

// MatchOption ...
type MatchOption func(r Request) (vars map[string]string, ok bool)

// Matches ...
type Matches []MatchOption

// Match ..
func (opts Matches) Match(r Request) (vars map[string]string, ok bool) {
	return Match(r)
}

// Match ...
func Match(r Request, opts ...MatchOption) (vars map[string]string, ok bool) {
	vars = make(map[string]string)
	for _, opt := range opts {
		res, ok := opt(r)
		if !ok {
			return nil, false
		}
		for k, v := range res {
			vars[k] = v
		}
	}
	return vars, true
}

// WithPath ...
func WithPath(path string) MatchOption {
	return func(r Request) (vars map[string]string, ok bool) {
		return nil, false
	}
}

// WithMethod ..
func WithMethod(method string) MatchOption {
	return func(r Request) (vars map[string]string, ok bool) {
		return nil, false
	}
}

// WithContentType ...
func WithContentType(typ string) MatchOption {
	return func(r Request) (vars map[string]string, ok bool) {
		return nil, false
	}
}

// WithContentJSON ..
func WithContentJSON() MatchOption {
	return WithContentType(ContentJSON)
}
