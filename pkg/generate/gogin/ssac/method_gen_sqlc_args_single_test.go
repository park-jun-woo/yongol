//ff:func feature=gen-gogin type=test control=sequence
//ff:what methodGen.sqlcArgsSingle 단위 테스트 (0개 → ctx, 1개 → ctx, <mapped>)
package ssac

import (
	"testing"
)

func TestMethodGenSqlcArgsSingle(t *testing.T) {
	t.Run("no inputs → ctx", func(t *testing.T) {
		g := &methodGen{}
		pre, args, imps := g.sqlcArgsSingle(map[string]string{})
		if pre != nil || args != "ctx" || imps != nil {
			t.Errorf("got (%v,%q,%v), want (nil,ctx,nil)", pre, args, imps)
		}
	})
	t.Run("single path-param input mapped", func(t *testing.T) {
		g := &methodGen{PathParams: map[string]bool{"id": true}}
		pre, args, imps := g.sqlcArgsSingle(map[string]string{"Id": "request.id"})
		if pre != nil {
			t.Errorf("unexpected preamble: %v", pre)
		}
		if args != "ctx, request.Id" {
			t.Errorf("args = %q, want %q", args, "ctx, request.Id")
		}
		if imps != nil {
			t.Errorf("no imports expected without pgtype col, got %v", imps)
		}
	})
}
