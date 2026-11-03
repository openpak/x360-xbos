// Command x360-xbos serves Xbox Live licensing (xbos), Live service 21 (see internal/xbos).
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/openpak/x360-common/httplog"
	"github.com/openpak/x360-common/wire"
	"github.com/openpak/x360-xbos/internal/xbos"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	addr := os.Getenv("X360_XBOS_ADDR")
	if addr == "" {
		addr = ":10021"
	}

	srv := &http.Server{Addr: addr, Handler: httplog.Wrap(wire.Guard(os.Getenv("X360_GATEWAY_KEY"), &xbos.Handler{Log: log})), ReadHeaderTimeout: 10 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		sc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(sc)
	}()
	log.Info("x360-xbos listening", "addr", addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("serve", "err", err)
		os.Exit(1)
	}
}
