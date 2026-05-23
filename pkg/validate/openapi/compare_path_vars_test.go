//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what comparePathVars — 일치/누락/초과 path 변수 진단 검증

package openapi

import (
	"testing"
)

func TestComparePathVars(t *testing.T) {
	tests := []struct {
		name      string
		want      map[string]bool // path template vars
		got       map[string]bool // declared params
		wantCount int
		wantSub   string
	}{
		{
			name:      "exact match returns nil",
			want:      map[string]bool{"id": true},
			got:       map[string]bool{"id": true},
			wantCount: 0,
		},
		{
			name:      "both empty returns nil",
			want:      map[string]bool{},
			got:       map[string]bool{},
			wantCount: 0,
		},
		{
			name:      "missing param in declarations",
			want:      map[string]bool{"id": true},
			got:       map[string]bool{},
			wantCount: 1,
			wantSub:   "template declares",
		},
		{
			name:      "extra param in declarations",
			want:      map[string]bool{},
			got:       map[string]bool{"extra": true},
			wantCount: 1,
			wantSub:   "parameters declares",
		},
		{
			name:      "both missing and extra",
			want:      map[string]bool{"id": true},
			got:       map[string]bool{"extra": true},
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := comparePathVars("/users/{id}", "get", 10, tt.want, tt.got)
			assertDiagCount(t, diags, tt.wantCount, tt.wantSub)
		})
	}
}
