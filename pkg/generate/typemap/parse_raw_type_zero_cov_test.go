//ff:func feature=gen-typemap type=test control=sequence
//ff:what typemap 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package typemap

import (
	"testing"
)

func TestParseRawType_ZeroCov(t *testing.T) {
	info := ParseRawType("VARCHAR(255)[]")
	if !info.IsArray || info.Param != "255" || info.Head != "VARCHAR" {
		t.Errorf("array/param/head wrong: %#v", info)
	}
	multi := ParseRawType("DOUBLE PRECISION")
	if !multi.MultiToken {
		t.Errorf("expected multi-token: %#v", multi)
	}
}
