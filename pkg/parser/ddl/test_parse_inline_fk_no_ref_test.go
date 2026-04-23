//ff:func feature=manifest type=test control=sequence
//ff:what parseInlineFK — REFERENCES 키워드 없으면 ok=false

package ddl

import "testing"

func TestParseInlineFK_NoRef(t *testing.T) {
	_, ok := parseInlineFK("user_id", []string{"user_id", "BIGINT", "NOT", "NULL"})
	if ok {
		t.Errorf("expected no inline FK")
	}
}
