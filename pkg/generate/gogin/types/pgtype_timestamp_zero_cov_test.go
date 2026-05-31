//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestPgtypeConstructorsZeroCov — 모든 pgtype 생성자 + unsupportedBinding 직접 커버
package types

import (
	"testing"
)

func TestPgtypeTimestamp_ZeroCov(t *testing.T) {
	cases := []struct {
		head string
		sqlc string
	}{
		{"TIMESTAMPTZ", "pgtype.Timestamptz"},
		{"TIMESTAMP", "pgtype.Timestamp"},
		{"DATE", "pgtype.Date"},
	}
	for _, c := range cases {
		if b := pgtypeTimestamp(c.head, true, ""); b.SqlcGoType != c.sqlc {
			t.Errorf("pgtypeTimestamp(%q) NOT NULL = %q, want %q", c.head, b.SqlcGoType, c.sqlc)
		}
		// nullable path appends Ptr to the convert funcs and uses *time.Time.
		if b := pgtypeTimestamp(c.head, false, ""); b.ApiField != "*time.Time" {
			t.Errorf("pgtypeTimestamp(%q) nullable api = %q", c.head, b.ApiField)
		}
	}
}
