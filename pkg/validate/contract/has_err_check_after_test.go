//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestHasErrCheckAfter — 블록 내 stmtIdx 의 err 를 뒤따르는 stmt 가 체크하는지 검증

package contract

import "testing"

func TestHasErrCheckAfter(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		idx     int
		errName string
		want    bool
	}{
		{
			name:    "checked after",
			body:    "err = f()\nif err != nil { return }\n",
			idx:     0,
			errName: "",
			want:    true,
		},
		{
			name:    "no check",
			body:    "err = f()\nreturn\n",
			idx:     0,
			errName: "",
			want:    false,
		},
		{
			name:    "clobbered before check",
			body:    "err = f()\nerr = g()\nif err != nil { return }\n",
			idx:     0,
			errName: "",
			want:    false,
		},
		{
			name:    "out of range",
			body:    "err = f()\n",
			idx:     5,
			errName: "",
			want:    false,
		},
		{
			name:    "negative idx",
			body:    "err = f()\n",
			idx:     -1,
			errName: "",
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmts := mustStmts(t, tt.body)
			if got := hasErrCheckAfter(stmts, tt.idx, tt.errName); got != tt.want {
				t.Fatalf("hasErrCheckAfter(idx=%d) = %v, want %v", tt.idx, got, tt.want)
			}
		})
	}
}
