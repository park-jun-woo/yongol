//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs19_NonDBBuiltin_NoDiagnostic — DB-free built-in 은 진단 없음

package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqs19_NonDBBuiltin_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "Hash", FileName: "svc.ssac",
			Sequences: []ssacparser.Sequence{{Type: "call", Package: "crypto", Model: "Hash"}},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"cache": cacheInterface()},
	}
	if diags := xqs19SsacBuiltinQueryRequired(fs); len(diags) != 0 {
		t.Errorf("crypto.Hash is not DB-using, expected no diag: %+v", diags)
	}
}
