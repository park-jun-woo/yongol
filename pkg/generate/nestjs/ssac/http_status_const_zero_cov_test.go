//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"strings"
	"testing"
)

func TestHttpStatusConst_ZeroCov(t *testing.T) {
	cases := map[int]string{
		400: "BAD_REQUEST", 401: "UNAUTHORIZED", 403: "FORBIDDEN",
		404: "NOT_FOUND", 409: "CONFLICT", 422: "UNPROCESSABLE_ENTITY",
		500: "INTERNAL_SERVER_ERROR",
	}
	for code, want := range cases {
		if got := httpStatusConst(code); got != want {
			t.Errorf("httpStatusConst(%d)=%q want %q", code, got, want)
		}
	}
	if got := httpStatusConst(418); !strings.Contains(got, "BAD_REQUEST") {
		t.Errorf("default branch: %q", got)
	}
}
