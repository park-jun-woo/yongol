//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-openapi
//ff:what resolveSeqErrStatus — 비guard/기본값/explicit ErrStatus/@call FuncSpec 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestResolveSeqErrStatus(t *testing.T) {
	specs := []funcspec.FuncSpec{
		{Package: "billing", Name: "Spend", ErrStatus: 402},
	}

	tests := []struct {
		name   string
		seq    ssac.Sequence
		want   int
		wantOK bool
	}{
		{
			name:   "non-guard type returns false",
			seq:    ssac.Sequence{Type: "response"},
			want:   0,
			wantOK: false,
		},
		{
			name:   "empty with default 404",
			seq:    ssac.Sequence{Type: "empty"},
			want:   404,
			wantOK: true,
		},
		{
			name:   "exists with default 409",
			seq:    ssac.Sequence{Type: "exists"},
			want:   409,
			wantOK: true,
		},
		{
			name:   "explicit ErrStatus overrides default",
			seq:    ssac.Sequence{Type: "empty", ErrStatus: 422},
			want:   422,
			wantOK: true,
		},
		{
			name:   "call with FuncSpec error status",
			seq:    ssac.Sequence{Type: "call", Model: "billing.Spend"},
			want:   402,
			wantOK: true,
		},
		{
			name:   "call without FuncSpec uses default 500",
			seq:    ssac.Sequence{Type: "call", Model: "other.Unknown"},
			want:   500,
			wantOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := resolveSeqErrStatus(tt.seq, specs)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if got != tt.want {
				t.Errorf("status = %d, want %d", got, tt.want)
			}
		})
	}
}
