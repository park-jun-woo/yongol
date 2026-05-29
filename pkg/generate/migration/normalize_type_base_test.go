//ff:func feature=migration type=test control=selection
//ff:what TestNormalizeTypeBase — base 토큰 정규화 + SERIAL 계열 감지
package migration

import "testing"

func TestNormalizeTypeBase(t *testing.T) {
	cases := []struct {
		upper      string
		wantBase   string
		wantSerial bool
	}{
		{"INT4", "INTEGER", false},
		{"INTEGER", "INTEGER", false},
		{"INT8", "BIGINT", false},
		{"BOOL", "BOOLEAN", false},
		{"CHARACTER VARYING", "VARCHAR", false},
		{"SERIAL", "INTEGER", true},
		{"BIGSERIAL", "BIGINT", true},
		{"SMALLSERIAL", "SMALLINT", true},
		{"CUSTOMTYPE", "CUSTOMTYPE", false},
	}
	for _, c := range cases {
		var ct CanonicalType
		gotSerial := normalizeTypeBase(c.upper, &ct)
		if ct.Base != c.wantBase || gotSerial != c.wantSerial {
			t.Errorf("normalizeTypeBase(%q) base=%q serial=%v, want base=%q serial=%v", c.upper, ct.Base, gotSerial, c.wantBase, c.wantSerial)
		}
	}
}
