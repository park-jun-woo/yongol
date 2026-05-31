//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderPrismaDataWhere(t *testing.T) {
	if got := renderPrismaData(nil); got != "{ data: body }" {
		t.Errorf("empty data = %q", got)
	}
	got := renderPrismaData([]ir.FieldArg{{Key: "title", ColumnName: "title", Literal: "x", IsQuoted: true}})
	if !strings.Contains(got, "data: { title: 'x' }") {
		t.Errorf("data = %q", got)
	}
	// all keys empty → fallback
	if got := renderPrismaData([]ir.FieldArg{{}}); got != "{ data: body }" {
		t.Errorf("empty-key fallback = %q", got)
	}

	if got := renderPrismaWhere(nil); got != "{}" {
		t.Errorf("empty where = %q", got)
	}
	w := renderPrismaWhere([]ir.FieldArg{{Key: "id", ColumnName: "id", Literal: "1"}})
	if !strings.Contains(w, "where: { id: 1 }") {
		t.Errorf("where = %q", w)
	}
}
