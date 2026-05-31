//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderSAData — FieldArg 배열 → SQLAlchemy keyword argument 문자열
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderSAData(t *testing.T) {
	t.Run("EmptyArgsSpreadsBody", func(t *testing.T) {
		if got := renderSAData(nil); got != "**body" {
			t.Errorf("got %q, want **body", got)
		}
	})
	t.Run("AllKeylessSpreadsBody", func(t *testing.T) {
		// FieldArgs with no ColumnName/Key/Field produce empty keys and are skipped.
		args := []ir.FieldArg{{Literal: "1"}, {Literal: "2"}}
		if got := renderSAData(args); got != "**body" {
			t.Errorf("got %q, want **body", got)
		}
	})
	t.Run("Kwargs", func(t *testing.T) {
		args := []ir.FieldArg{
			{ColumnName: "title", Location: ir.LocBody},
			{Key: "OrgID", Location: ir.LocUser, ColumnName: "org_id"},
		}
		got := renderSAData(args)
		want := "title=body.title, org_id=current_user[\"org_id\"]"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
