//ff:func feature=cli-init type=test control=iteration dimension=1
//ff:what TestParseHTTPPath — 정상 메서드 / 필드부족 / 잘못된 메서드 에러 분기 검증
package cliinit

import (
	"testing"
)

func TestParseHTTPPath_Valid(t *testing.T) {
	for _, m := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		r, err := parseHTTPPath(m + " /tasks/{id}")
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", m, err)
		}
		if r.URI != "/tasks/{id}" {
			t.Errorf("%s: URI = %q", m, r.URI)
		}
	}
}
