//ff:func feature=manifest type=test control=sequence
//ff:what parseConstraintFK — FK 패턴 불일치 라인은 ok=false

package ddl

import "testing"

func TestParseConstraintFK_Invalid(t *testing.T) {
	if _, ok := parseConstraintFK("not a fk line"); ok {
		t.Errorf("expected false")
	}
}
