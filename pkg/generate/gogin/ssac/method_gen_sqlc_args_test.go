//ff:func feature=gen-gogin type=test control=sequence
//ff:what methodGen.sqlcArgs 단위 테스트 (0/1/N 입력 분기 디스패치 + activeMethod 정리)
package ssac

import (
	"testing"
)

func TestMethodGenSqlcArgs(t *testing.T) {
	g := &methodGen{PathParams: map[string]bool{"id": true}}

	t.Run("zero inputs → ctx", func(t *testing.T) {
		_, args, _ := g.sqlcArgs("WorkflowList", map[string]string{})
		if args != "ctx" {
			t.Errorf("args = %q, want ctx", args)
		}
		if g.activeMethod != "" {
			t.Errorf("activeMethod not cleared: %q", g.activeMethod)
		}
	})

	t.Run("single input dispatches to single form", func(t *testing.T) {
		_, args, _ := g.sqlcArgs("WorkflowFindByID", map[string]string{"Id": "request.id"})
		if args != "ctx, request.Id" {
			t.Errorf("args = %q, want %q", args, "ctx, request.Id")
		}
	})

	t.Run("multi input dispatches to params struct", func(t *testing.T) {
		_, args, _ := g.sqlcArgs("WorkflowUpdate", map[string]string{
			"Id":     "request.id",
			"Status": `"open"`,
		})
		want := `ctx, db.WorkflowUpdateParams{Id: request.Id, Status: "open"}`
		if args != want {
			t.Errorf("args = %q, want %q", args, want)
		}
	})
}
