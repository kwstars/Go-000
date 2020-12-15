package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ListenSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGUSR1)
	log.Printf("Start ListenSignal")
	for {
		select {
		case <-ctx.Done():
			log.Println("Stop ListenSignal")
			return nil
		case s := <-c:
			log.Printf("Get signal %v", s)
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, os.Interrupt:
				return errors.New("Gracefully shutdown the server")
			case syscall.SIGHUP:
				fmt.Println("Reload configuration file")
			case syscall.SIGUSR1:
				fmt.Println("Switch log mode")
			}
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
	g, ctx := errgroup.WithContext(context.Background())

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
		return ListenSignal(ctx)
	})

	// inject error
	g.Go(func() error {
		fmt.Println("inject")
		time.Sleep(time.Second)
		fmt.Println("inject finish")
		return errors.New("inject error")
	})

	err := g.Wait()
	log.Printf("%v\n", err)
}
