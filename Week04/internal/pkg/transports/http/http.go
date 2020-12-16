package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

var ProviderSet = wire.NewSet(InitRouter, New)

type InitControllers func(r *gin.Engine)

type Server struct {
	router     *gin.Engine
	httpServer http.Server
}

func InitRouter(init InitControllers) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery()) // panic之后自动恢复
	init(r)

	return r
}

func New(r *gin.Engine) (*Server, error) {
	var s = &Server{
		router: r,
	}

	return s, nil
}

func (s *Server) Start() error {
	s.httpServer = http.Server{Addr: "127.0.0.1:44444", Handler: s.router}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			return
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	log.Println("Stop server")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Shutdown http server err")
	}

	return nil
}
