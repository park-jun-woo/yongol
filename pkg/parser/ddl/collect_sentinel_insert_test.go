//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what collectSentinelInsert — INSERT 본문을 종결 `;` 까지 수집 (단일/다중 라인, 인용 인식)

package ddl

import "testing"

func TestCollectSentinelInsert(t *testing.T) {
	t.Run("single line insert", func(t *testing.T) {
		lines := []string{"INSERT INTO roles VALUES (1, 'a'); extra"}
		res, next := collectSentinelInsert(lines, 0, "roles", true)
		if res.SQL != "INSERT INTO roles VALUES (1, 'a');" {
			t.Errorf("SQL = %q", res.SQL)
		}
		if res.Table != "roles" || !res.Annotated || res.StartLine != 1 {
			t.Errorf("res = %+v", res)
		}
		if next != 1 {
			t.Errorf("next = %d, want 1", next)
		}
	})
	t.Run("multi line insert", func(t *testing.T) {
		lines := []string{"INSERT INTO t (id, name)", "VALUES", "(1, 'x');"}
		res, next := collectSentinelInsert(lines, 0, "t", false)
		want := "INSERT INTO t (id, name)\nVALUES\n(1, 'x');"
		if res.SQL != want {
			t.Errorf("SQL = %q, want %q", res.SQL, want)
		}
		if next != 3 {
			t.Errorf("next = %d, want 3", next)
		}
	})
	t.Run("semicolon inside literal not terminator", func(t *testing.T) {
		lines := []string{"INSERT INTO t VALUES ('a;b'", "); done"}
		res, _ := collectSentinelInsert(lines, 0, "t", false)
		want := "INSERT INTO t VALUES ('a;b'\n);"
		if res.SQL != want {
			t.Errorf("SQL = %q, want %q", res.SQL, want)
		}
	})
}
