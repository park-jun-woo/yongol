//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderResponseOpSourceCasing -- tsSourceExpr 으로 PascalCase → camelCase 변환 검증
package ssac

import (
	"testing"
)

func TestTsSourceExpr(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "PascalCase field to camelCase",
			source: "token.AccessToken",
			want:   "token.accessToken",
		},
		{
			name:   "already lowercase",
			source: "token.email",
			want:   "token.email",
		},
		{
			name:   "no dot — bare variable",
			source: "user",
			want:   "user",
		},
		{
			name:   "multi-word PascalCase",
			source: "token.RefreshToken",
			want:   "token.refreshToken",
		},
		{
			name:   "single char field",
			source: "x.Y",
			want:   "x.y",
		},
		{
			name:   "empty string",
			source: "",
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tsSourceExpr(tt.source)
			if got != tt.want {
				t.Errorf("tsSourceExpr(%q) = %q, want %q", tt.source, got, tt.want)
			}
		})
	}
}
