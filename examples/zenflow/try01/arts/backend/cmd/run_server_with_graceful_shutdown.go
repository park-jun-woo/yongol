//ff:func feature=main type=util control=sequence
//ff:what runServerWithGracefulShutdown — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=847be61d
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func runServerWithGracefulShutdown(r *gin.Engine, cancelBootstrap context.CancelFunc, headerLimit int) {
	httpSrv := &http.Server{Addr: ":8080", Handler: r, MaxHeaderBytes: headerLimit}
	go func() {
		slog.Info("server starting", "addr", ":8080")
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server", "err", err)
			os.Exit(1)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("shutting down")
	cancelBootstrap()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown", "err", err)
	}
	slog.Info("shutdown complete")
}
