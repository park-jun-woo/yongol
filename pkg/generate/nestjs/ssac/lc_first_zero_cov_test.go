//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"
)

func TestLcFirst_ZeroCov(t *testing.T) {
	if lcFirst("Hello") != "hello" {
		t.Error("lcFirst failed")
	}
	if lcFirst("") != "" {
		t.Error("empty failed")
	}
}
