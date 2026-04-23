//ff:func feature=manifest type=test control=sequence
//ff:what parseCheckEnum — CHECK 절이 없으면 빈 값 반환

package ddl

import "testing"

func TestParseCheckEnum_NoMatch(t *testing.T) {
	col, vals := parseCheckEnum("no check clause here")
	if col != "" || vals != nil {
		t.Errorf("got col=%q vals=%v, want empty", col, vals)
	}
}
