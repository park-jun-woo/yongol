//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs19_Cache_Present_NoDiagnostic — cache @call 쿼리 존재 시 진단 없음

package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqs19_Cache_Present_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "PutEntry", FileName: "svc.ssac",
			Sequences: []ssacparser.Sequence{{Type: "call", Package: "cache", Model: "Set"}},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"cache": cacheInterface()},
		SQLcQueries:    []sqlcparser.QuerySpec{{Name: "CacheSet"}},
	}
	if diags := xqs19SsacBuiltinQueryRequired(fs); len(diags) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", diags)
	}
}
