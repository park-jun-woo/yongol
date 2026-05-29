//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what neutralCode 단위 테스트 (status → 기본 에러 코드, unknown fallback)

package ssac

import "testing"

func TestNeutralCode(t *testing.T) {
	cases := map[int]string{
		400: "bad_request",
		401: "unauthorized",
		402: "payment_required",
		404: "not_found",
		429: "too_many_requests",
		500: "internal_error",
		418: "internal_error", // unknown → fallback
	}
	for in, want := range cases {
		if got := neutralCode(in); got != want {
			t.Errorf("neutralCode(%d) = %q, want %q", in, got, want)
		}
	}
}
