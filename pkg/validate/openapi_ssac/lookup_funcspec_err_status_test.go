//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-openapi
//ff:what lookupFuncSpecErrStatus — 미매칭/매칭/단일 파트 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestLookupFuncSpecErrStatus(t *testing.T) {
	specs := []funcspec.FuncSpec{
		{Package: "billing", Name: "Spend", ErrStatus: 402},
		{Package: "billing", Name: "spend", ErrStatus: 402},
		{Package: "auth", Name: "VerifyPassword", ErrStatus: 401},
	}

	tests := []struct {
		name  string
		model string
		want  int
	}{
		{"matching exact name", "billing.Spend", 402},
		{"matching lowercase name", "billing.spend", 402},
		{"matching other package", "auth.VerifyPassword", 401},
		{"no matching func", "billing.Unknown", 0},
		{"no dot returns 0", "billing", 0},
		{"empty model", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lookupFuncSpecErrStatus(tt.model, specs)
			if got != tt.want {
				t.Errorf("lookupFuncSpecErrStatus(%q) = %d, want %d", tt.model, got, tt.want)
			}
		})
	}
}
