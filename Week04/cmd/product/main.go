package main

import (
	"fmt"
	"github.com/kwstars/Go-000/Week04/internal/app/product/di"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	err = app.Start()
	if err != nil {
		fmt.Println(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Println("app exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
			//dosomething
		default:
			return
		}
	}
}
