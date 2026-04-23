//ff:func feature=manifest type=test control=sequence
//ff:what parseInlineFK — REFERENCES users(id) 인라인 FK 추출

package ddl

import "testing"

func TestParseInlineFK_Match(t *testing.T) {
	fk, ok := parseInlineFK("user_id", []string{"user_id", "BIGINT", "NOT", "NULL", "REFERENCES", "users(id),"})
	if !ok {
		t.Fatalf("expected ok")
	}
	if fk.Column != "user_id" || fk.RefTable != "users" || fk.RefColumn != "id" {
		t.Errorf("fk = %+v", fk)
	}
}
