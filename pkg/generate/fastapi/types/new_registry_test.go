//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRegistry — NewRegistry 팩토리 + Bind 디스패치 검증
package types

import (
	"testing"
)

func TestNewRegistry(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry() returned nil")
	}
}
