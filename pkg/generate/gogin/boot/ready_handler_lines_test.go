//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestReadyHandlerLines — /ready 핸들러 DB 유무 분기 검증

package boot

import (
	"strings"
	"testing"
)

func TestReadyHandlerLines(t *testing.T) {
	t.Run("WithoutDB", func(t *testing.T) {
		lines := readyHandlerLines(false)
		body := strings.Join(lines, "\n")
		if !strings.Contains(body, `r.GET("/ready"`) {
			t.Errorf("expected /ready route, got:\n%s", body)
		}
		if !strings.Contains(body, `c.JSON(200, gin.H{"status": "ok"})`) {
			t.Errorf("expected static 200 handler, got:\n%s", body)
		}
		if strings.Contains(body, "readyHandlerWithDB") {
			t.Errorf("should not reference DB helper without DB, got:\n%s", body)
		}
	})

	t.Run("WithDB", func(t *testing.T) {
		lines := readyHandlerLines(true)
		body := strings.Join(lines, "\n")
		if !strings.Contains(body, "readyHandlerWithDB(pool)") {
			t.Errorf("expected DB helper delegation, got:\n%s", body)
		}
	})
}
