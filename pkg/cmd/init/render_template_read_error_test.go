//ff:func feature=cli-init type=test control=sequence
//ff:what TestRenderTemplate — embed 읽기 에러 / 정상 렌더(데이터 치환) 분기 검증
package cliinit

import (
	"testing"
)

func TestRenderTemplate_ReadError(t *testing.T) {
	if _, err := renderTemplate("templates/missing.tmpl", templateData{}); err == nil {
		t.Fatal("want error for missing embed file")
	}
}
