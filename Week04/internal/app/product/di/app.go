package di

import (
	"github.com/kwstars/Go-000/Week04/internal/app/product/services"
	"github.com/kwstars/Go-000/Week04/internal/pkg/transports/http"
	"github.com/pkg/errors"
	"log"
)

type App struct {
	svc        *services.ProductService
	httpServer *http.Server
}

func NewApp(svc *services.ProductService, h *http.Server) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:        svc,
		httpServer: h,
	}
	closeFunc = func() {
		if err := h.Stop(); err != nil {
			log.Println("httpSrv.Shutdown error(%v)", err)
		}
	}
	return
}

func (a *App) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}
	return nil
}
