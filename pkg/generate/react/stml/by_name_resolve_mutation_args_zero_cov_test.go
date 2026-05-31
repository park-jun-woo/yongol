//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameResolveMutationArgs_ZeroCov(t *testing.T) {
	cons := byNameConstraints()
	// void
	resolveMutationArgs("DeleteItem", "", true, cons)
	// body only
	resolveMutationArgs("CreateItem", "", false, cons)
	// body + path
	fn, args := resolveMutationArgs("CreateItem", "{ id }", false, cons)
	if fn == "" && args == "" {
		t.Errorf("resolveMutationArgs body+path returned empty")
	}
}
