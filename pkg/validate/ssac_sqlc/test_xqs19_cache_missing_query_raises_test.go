//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs19_Cache_Missing_Query_Raises — cache @call 쿼리 부재 → [XQS-19]

package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqs19_Cache_Missing_Query_Raises(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "PutEntry", FileName: "svc.ssac",
			Sequences: []ssacparser.Sequence{{Type: "call", Package: "cache", Model: "Set"}},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"cache": cacheInterface()},
	}
	diags := xqs19SsacBuiltinQueryRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "CacheSet") {
		t.Errorf("diag missing CacheSet: %s", diags[0].Message)
	}
}
