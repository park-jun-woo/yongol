//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what resolveErrStatus 단위 테스트 (explicit > 타입 기본값)

package ssac

import "testing"

func TestResolveErrStatus(t *testing.T) {
	cases := []struct {
		name     string
		seqType  string
		explicit int
		want     int
	}{
		{"explicit overrides", "empty", 422, 422},
		{"empty default 404", "empty", 0, 404},
		{"exists default 409", "exists", 0, 409},
		{"state default 409", "state", 0, 409},
		{"auth default 403", "auth", 0, 403},
		{"call default 500", "call", 0, 500},
		{"eval fallback 500", "eval", 0, 500},
		{"unknown default 500", "mystery", 0, 500},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := resolveErrStatus(tc.seqType, tc.explicit); got != tc.want {
				t.Errorf("resolveErrStatus(%q,%d) = %d, want %d", tc.seqType, tc.explicit, got, tc.want)
			}
		})
	}
}
