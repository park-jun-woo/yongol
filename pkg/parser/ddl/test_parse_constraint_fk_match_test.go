//ff:func feature=manifest type=test control=sequence
//ff:what parseConstraintFK — CONSTRAINT fk_... FOREIGN KEY ... 라인 파싱

package ddl

import "testing"

func TestParseConstraintFK_Match(t *testing.T) {
	fk, ok := parseConstraintFK("CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)")
	if !ok {
		t.Fatalf("expected ok")
	}
	if fk.Column != "user_id" || fk.RefTable != "users" || fk.RefColumn != "id" {
		t.Errorf("fk = %+v", fk)
	}
}
