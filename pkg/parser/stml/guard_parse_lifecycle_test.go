//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what parseGuardLifecycle — ref "." lifecycle 파싱 (loading/error/empty, 비식별자/잘못된 키워드 에러) 검증

package stml

import "testing"

func TestParseGuardLifecycle(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantErr       bool
		wantLifecycle string
	}{
		{name: "loading", input: "a.b.loading", wantLifecycle: "loading"},
		{name: "error", input: "a.b.error", wantLifecycle: "error"},
		{name: "empty", input: "a.b.empty", wantLifecycle: "empty"},
		{name: "non-ident after dot", input: "a.b.(", wantErr: true},
		{name: "invalid keyword", input: "a.b.pending", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertParseGuardLifecycle(t, tt.input, tt.wantErr, tt.wantLifecycle)
		})
	}
}
