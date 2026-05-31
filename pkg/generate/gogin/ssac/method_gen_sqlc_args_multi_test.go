//ff:func feature=gen-gogin type=test control=sequence
//ff:what methodGen.sqlcArgsMulti 단위 테스트 (정렬된 Params 구조체 리터럴 방출)
package ssac

import (
	"testing"
)

func TestMethodGenSqlcArgsMulti(t *testing.T) {
	g := &methodGen{
		PathParams: map[string]bool{"id": true},
	}
	inputs := map[string]string{
		"Status": `"open"`,
		"Id":     "request.id",
	}
	pre, args, _ := g.sqlcArgsMulti("WorkflowUpdate", inputs)
	if pre != nil {
		t.Errorf("unexpected preamble: %v", pre)
	}
	// Fields are sorted: "Id: ..." before "Status: ...".
	want := `ctx, db.WorkflowUpdateParams{Id: request.Id, Status: "open"}`
	if args != want {
		t.Errorf("args = %q, want %q", args, want)
	}
}
