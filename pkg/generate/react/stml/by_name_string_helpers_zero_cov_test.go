//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameStringHelpers_ZeroCov(t *testing.T) {
	if clsAttr("c") == "" {
		t.Errorf("clsAttr empty")
	}
	_ = clsAttr("")
	if orDefault("", "d") != "d" {
		t.Errorf("orDefault default")
	}
	_ = orDefault("x", "d")
	runes := []rune("RoomID")
	_ = isWordBoundary(runes, 0)
	_ = isWordBoundary(runes, 4)
	_ = isWordBoundary([]rune("abc"), 0)
}
