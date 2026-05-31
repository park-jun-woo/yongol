//ff:func feature=cli-init type=test control=sequence
//ff:what TestParseHTTPPath — 정상 메서드 / 필드부족 / 잘못된 메서드 에러 분기 검증
package cliinit

import (
	"testing"
)

func TestParseHTTPPath_FieldCount(t *testing.T) {
	if _, err := parseHTTPPath("badpath"); err == nil {
		t.Fatal("want error for missing method/uri")
	}
	if _, err := parseHTTPPath("GET /a /b"); err == nil {
		t.Fatal("want error for too many fields")
	}
}
