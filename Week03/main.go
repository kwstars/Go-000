package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ListenSingle(stopFunc func()) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	for {
		s := <-c
		log.Printf("Get signal %v", s)
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, os.Interrupt:
			log.Println("App exit")
			stopFunc()
			time.Sleep(3 * time.Second)
			return nil
		case syscall.SIGHUP:
		default:
			return nil
		}
	}
}

func WebService(ctx context.Context, name, port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "okay") })

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		shutdownctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		log.Printf("Shutdown " + name)
		srv.Shutdown(shutdownctx)
	}()

	log.Println("Start " + name)
	return srv.ListenAndServe()
}

func main() {
	cancel, cancelFunc := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(cancel)

	//server1
	g.Go(func() error {
		return WebService(ctx, "server1", "8081")
	})

	//server2
	g.Go(func() error {
		return WebService(ctx, "server2", "8082")
	})

	//接收信号
	g.Go(func() error {
		return ListenSingle(cancelFunc)
	})

	if err := g.Wait(); err != nil {
		log.Printf("Server Error:%v\n", err)
	}
}
