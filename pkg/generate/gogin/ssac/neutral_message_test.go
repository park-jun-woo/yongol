//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what neutralMessage 단위 테스트 (status → client-safe 메시지, unknown fallback)

package ssac

import "testing"

func TestNeutralMessage(t *testing.T) {
	cases := map[int]string{
		400: "Bad request",
		401: "Unauthorized",
		402: "Payment required",
		404: "Not found",
		500: "Internal error",
		418: "Internal error", // unknown → fallback
	}
	for in, want := range cases {
		if got := neutralMessage(in); got != want {
			t.Errorf("neutralMessage(%d) = %q, want %q", in, got, want)
		}
	}
}
