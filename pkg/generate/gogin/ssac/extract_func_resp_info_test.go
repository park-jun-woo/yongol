//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what extractFuncRespInfo 단위 테스트 (Model alias + import path 추출)
package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestExtractFuncRespInfo(t *testing.T) {
	imports := []string{
		"github.com/park-jun-woo/zenflow/internal/dashboard",
		"github.com/park-jun-woo/zenflow/internal/other",
	}
	t.Run("matched alias resolves import path", func(t *testing.T) {
		seq := ssacparser.Sequence{Model: "dashboard.Summarize"}
		got := extractFuncRespInfo(seq, imports)
		if got.PkgAlias != "dashboard" {
			t.Errorf("PkgAlias = %q, want dashboard", got.PkgAlias)
		}
		if got.ImportPath != "github.com/park-jun-woo/zenflow/internal/dashboard" {
			t.Errorf("ImportPath = %q", got.ImportPath)
		}
	})
	t.Run("alias without matching import → empty path", func(t *testing.T) {
		seq := ssacparser.Sequence{Model: "unknown.Func"}
		got := extractFuncRespInfo(seq, imports)
		if got.PkgAlias != "unknown" || got.ImportPath != "" {
			t.Errorf("got %+v, want alias=unknown path=\"\"", got)
		}
	})
	t.Run("model without dot → empty alias", func(t *testing.T) {
		got := extractFuncRespInfo(ssacparser.Sequence{Model: "Bare"}, imports)
		if got.PkgAlias != "" {
			t.Errorf("PkgAlias = %q, want empty", got.PkgAlias)
		}
	})
}
