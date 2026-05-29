//ff:func feature=gen-gogin type=test control=sequence
//ff:what ginRunHelperServerLifecycle — http.Server 기동 + SIGINT/SIGTERM graceful shutdown 헬퍼 소스 반환

package boot

import (
	"strings"
	"testing"
)

func TestGinRunHelperServerLifecycle(t *testing.T) {
	src := ginRunHelperServerLifecycle()
	for _, must := range []string{
		"func runServerWithGracefulShutdown(r *gin.Engine, cancelBootstrap context.CancelFunc, headerLimit int) {",
		"MaxHeaderBytes: headerLimit",
		"httpSrv.ListenAndServe()",
		"signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)",
		"cancelBootstrap()",
		"context.WithTimeout(context.Background(), 10*time.Second)",
		"httpSrv.Shutdown(shutdownCtx)",
	} {
		if !strings.Contains(src, must) {
			t.Errorf("server lifecycle helper missing %q, got:\n%s", must, src)
		}
	}
}
