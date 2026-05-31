//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestStructNameForTable — exported wrapper가 structNameFor와 동일 결과 반환 검증
package sqlcpost

import (
	"testing"
)

func TestStructNameForTable(t *testing.T) {
	cases := []string{"users", "audit_logs", "categories", "data"}
	for _, c := range cases {
		if got, want := StructNameForTable(c), structNameFor(c); got != want {
			t.Errorf("StructNameForTable(%q) = %q, want %q", c, got, want)
		}
	}
	if got := StructNameForTable("users"); got != "User" {
		t.Errorf("StructNameForTable(\"users\") = %q, want User", got)
	}
}
