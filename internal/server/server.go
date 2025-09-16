package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

func Start(ctx context.Context) error {
	e := echo.New()
	err := routes(e)
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	go func() {
		cfg := do.MustInvoke[*config.Config](bootstrap.Injector)
		err := e.Start(fmt.Sprintf(":%d", cfg.Server.Port))
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("failed to start server:", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	return e.Shutdown(ctx)
}
