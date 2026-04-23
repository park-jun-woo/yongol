//ff:func feature=manifest type=test control=sequence
//ff:what parseCheckEnum — CHECK(IN (...)) 절에서 column 이름 + enum 값 3개 추출

package ddl

import "testing"

func TestParseCheckEnum_Basic(t *testing.T) {
	col, vals := parseCheckEnum(`status VARCHAR(32) NOT NULL CHECK (status IN ('a','b','c'))`)
	if col != "status" {
		t.Errorf("col = %q, want status", col)
	}
	if len(vals) != 3 {
		t.Fatalf("vals = %v", vals)
	}
}
