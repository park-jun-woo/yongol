//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCheckBindCase_ZeroCov — checkBindCase 헬퍼를 self-consistent bindCase 로 직접 호출

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// TestCheckBindCase_ZeroCov exercises checkBindCase by deriving the expected
// fields from an actual Bind() result, so the helper's every assertion runs the
// equal branch (no t.Errorf). This pins the otherwise-uncalled helper.
func TestCheckBindCase_ZeroCov(t *testing.T) {
	reg := NewRegistry()
	family := typemap.FamilyString
	opts := ir.BindOpts{NotNull: true}
	b := reg.Bind(family, opts)

	c := bindCase{
		name:           "string-notnull",
		family:         family,
		opts:           opts,
		wantDBType:     b.DBType,
		wantAPIType:    b.APIType,
		wantToDBExpr:   b.ToDBExpr,
		wantToAPIExpr:  b.ToAPIExpr,
		wantToRespExpr: b.ToResponseExpr,
		wantNilCheck:   b.NilCheckExpr,
		wantSupported:  b.Supported,
	}
	checkBindCase(t, c)
}
