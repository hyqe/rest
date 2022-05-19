package rest

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"context"
)

func Runf(handler http.HandlerFunc, opts ...SeverOption) error {
	return Run(handler, opts...)
}

func Run(handler http.Handler, opts ...SeverOption) error {
	return NewServer().
		Apply(defaultServerOptions()...).
		Apply(WithHandler(handler)).
		Apply(opts...).
		Run()
}

type Server struct {
	ip         string
	port       int
	handler    http.Handler
	purge      time.Duration
	interrupts []os.Signal
	cert       string
	key        string
}

func NewServer() *Server {
	return &Server{}
}

func defaultServerOptions() []SeverOption {
	return []SeverOption{
		WithPort(8080),
		WithPurge(time.Second * 3),
		WithInterrupts(defaultInterruptSignals...),
		WithHandler(NewStatusNotImplementedHandler()),
	}
}

func (s *Server) Apply(opts ...SeverOption) *Server {
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%v:%v", s.ip, s.port)
}

func (s *Server) Handler() http.Handler {
	return s.handler
}

func (s *Server) SecureKeys() (cert, key string, ok bool) {
	if s.cert != "" && s.key != "" {
		return s.cert, s.key, true
	}
	return cert, key, false
}

func (s *Server) Run() error {
	httpSrv := http.Server{
		Addr:    s.Addr(),
		Handler: s.Handler(),
	}

	go func() {
		<-Interrupt(s.interrupts...)
		ctx, cancel := context.WithTimeout(context.Background(), s.purge)
		defer cancel()
		httpSrv.Shutdown(ctx)
	}()

	if cert, key, ok := s.SecureKeys(); ok {
		return httpSrv.ListenAndServeTLS(cert, key)
	} else {
		return httpSrv.ListenAndServe()
	}
}

type SeverOption func(*Server)

func WithIP(ip string) SeverOption {
	return func(s *Server) {
		s.ip = ip
	}
}

func WithPort(port int) SeverOption {
	return func(s *Server) {
		s.port = port
	}
}

func WithHandler(h http.Handler) SeverOption {
	return func(s *Server) {
		s.handler = h
	}
}

func WithPurge(d time.Duration) SeverOption {
	return func(s *Server) {
		s.purge = d
	}
}

func WithInterrupts(sig ...os.Signal) SeverOption {
	return func(s *Server) {
		s.interrupts = sig
	}
}

func WithTLS(cert, key string) SeverOption {
	return func(s *Server) {
		s.cert = cert
		s.key = key
	}
}
