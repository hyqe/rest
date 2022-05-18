package rest

import (
	"os"
	"os/signal"
	"syscall"
)

func Interrupt(s ...os.Signal) <-chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, s...)
	return done
}

var defaultInterruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGINT,
	syscall.SIGTERM,
}
