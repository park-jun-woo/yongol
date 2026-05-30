//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestNullableAPIType — NotNull 여부에 따른 " | None" 접미사 검증

package types

import "testing"

func TestNullableAPIType(t *testing.T) {
	if got := nullableAPIType("int", true); got != "int" {
		t.Errorf("nullableAPIType(int, true) = %q, want int", got)
	}
	if got := nullableAPIType("int", false); got != "int | None" {
		t.Errorf("nullableAPIType(int, false) = %q, want int | None", got)
	}
}
