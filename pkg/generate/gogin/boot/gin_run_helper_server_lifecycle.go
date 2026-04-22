//ff:func feature=gen-gogin type=generator control=sequence
//ff:what ginRunHelperServerLifecycle — http.Server 기동 + SIGINT/SIGTERM graceful shutdown 헬퍼 소스 반환

package boot

// ginRunHelperServerLifecycle returns the top-level
// runServerWithGracefulShutdown(r *gin.Engine, cancelBootstrap context.CancelFunc, headerLimit int)
// helper source. Extracted from main() to keep main under filefunc Q3's 100-line
// limit for control=sequence funcs. The helper launches http.Server in a goroutine,
// waits on SIGINT/SIGTERM, cancels the bootstrap context, then performs a 10s
// graceful Shutdown. headerLimit is wired into http.Server.MaxHeaderBytes so
// oversized request-line / header payloads return 431 at the stdlib layer
// before reaching any gin handler.
func ginRunHelperServerLifecycle() string {
	return `func runServerWithGracefulShutdown(r *gin.Engine, cancelBootstrap context.CancelFunc, headerLimit int) {
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
}`
}
