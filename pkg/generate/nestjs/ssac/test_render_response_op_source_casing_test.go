//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderResponseOpSourceCasing -- tsSourceExpr 으로 PascalCase → camelCase 변환 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
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

func TestRenderResponseOpSourceCasing(t *testing.T) {
	op := &ir.ResponseOp{
		Fields: []ir.ResponseField{
			{Name: "access_token", Source: "token.AccessToken"},
			{Name: "refresh_token", Source: "token.RefreshToken"},
			{Name: "email", Source: "user.Email"},
		},
	}

	var b strings.Builder
	renderResponseOp(&b, op, "    ")
	got := b.String()

	// PascalCase fields should be converted to camelCase.
	if !strings.Contains(got, "token.accessToken") {
		t.Errorf("expected token.accessToken in output, got:\n%s", got)
	}
	if !strings.Contains(got, "token.refreshToken") {
		t.Errorf("expected token.refreshToken in output, got:\n%s", got)
	}
	if !strings.Contains(got, "user.email") {
		t.Errorf("expected user.email in output, got:\n%s", got)
	}

	// PascalCase should NOT appear in output.
	if strings.Contains(got, "AccessToken") {
		t.Errorf("PascalCase AccessToken should not appear in output, got:\n%s", got)
	}
}

func TestRenderResponseOpSingleVar(t *testing.T) {
	op := &ir.ResponseOp{
		SingleVar: "course",
	}

	var b strings.Builder
	renderResponseOp(&b, op, "    ")
	got := b.String()

	want := "    return course;\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
