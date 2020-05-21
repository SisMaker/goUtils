package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGUSR1)
	signal.Notify(ch)

	for {
		sig := <-ch
		switch sig {
		case syscall.SIGHUP:
			println("SIGSTOP")
			return
		case syscall.SIGQUIT:
			println("SIGQUIT")
			return
		case syscall.SIGTERM:
			println("SIGTERM")
			return
		case syscall.SIGSTOP:
			println("SIGSTOP")
			return
		case syscall.SIGUSR1:
			println("SIGUSR1")
			return
		default:
			println("what is this ", sig)
			return
		}
	}
}
