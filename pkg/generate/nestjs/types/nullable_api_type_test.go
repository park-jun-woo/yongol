//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNullableAPIType — NotNull 분기 커버
package types

import "testing"

func TestNullableAPIType_ZeroCov(t *testing.T) {
	if got := nullableAPIType("number", true); got != "number" {
		t.Errorf("NOT NULL = %q", got)
	}
	if got := nullableAPIType("number", false); got != "number | null" {
		t.Errorf("nullable = %q", got)
	}
}
