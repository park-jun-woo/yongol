//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what healthHelperReadyWithDB — /ready pgxpool ping 핸들러 생성 헬퍼 소스 반환
package boot

import (
	"strings"
	"testing"
)

func TestHealthHelperReadyWithDB(t *testing.T) {
	src := healthHelperReadyWithDB()
	for _, must := range []string{
		"func readyHandlerWithDB(pool *pgxpool.Pool) gin.HandlerFunc {",
		"context.WithTimeout(c.Request.Context(), 2*time.Second)",
		"pool.Ping(pingCtx)",
		"c.JSON(503,",
		"c.JSON(200,",
	} {
		if !strings.Contains(src, must) {
			t.Errorf("readyHandlerWithDB helper missing %q, got:\n%s", must, src)
		}
	}
	if strings.Contains(src, "sql.DB") || strings.Contains(src, "conn.Ping") {
		t.Errorf("helper must use pgxpool, not database/sql, got:\n%s", src)
	}
}
