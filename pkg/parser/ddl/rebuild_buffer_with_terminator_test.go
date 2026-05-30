//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what rebuildBufferWithTerminator — i==j 단일 라인 절단 / 다중 라인 연결

package ddl

import (
	"strings"
	"testing"
)

func TestRebuildBufferWithTerminator(t *testing.T) {
	t.Run("terminator on insert line itself", func(t *testing.T) {
		var buf strings.Builder
		buf.WriteString("garbage")
		lines := []string{"INSERT INTO t VALUES (1); -- trailing"}
		k := strings.Index(lines[0], ";")
		rebuildBufferWithTerminator(&buf, lines, 0, 0, lines[0], k)
		want := "INSERT INTO t VALUES (1);"
		if buf.String() != want {
			t.Errorf("buf = %q, want %q", buf.String(), want)
		}
	})
	t.Run("terminator on later line", func(t *testing.T) {
		var buf strings.Builder
		buf.WriteString("old")
		lines := []string{"INSERT INTO t", "VALUES", "(1); rest"}
		ln := lines[2]
		k := strings.Index(ln, ";")
		rebuildBufferWithTerminator(&buf, lines, 0, 2, ln, k)
		want := "INSERT INTO t\nVALUES\n(1);"
		if buf.String() != want {
			t.Errorf("buf = %q, want %q", buf.String(), want)
		}
	})
}
