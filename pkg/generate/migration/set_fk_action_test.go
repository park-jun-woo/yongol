//ff:func feature=migration type=test control=iteration dimension=1
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestSetFKAction(t *testing.T) {
	cases := []struct {
		action, val            string
		wantDelete, wantUpdate string
	}{
		{"DELETE", "CASCADE", "CASCADE", ""},
		{"UPDATE", "SET NULL", "", "SET NULL"},
		{"OTHER", "CASCADE", "", ""},
	}
	for _, c := range cases {
		c := c
		t.Run(c.action, func(t *testing.T) {
			fk := &ForeignKey{}
			setFKAction(fk, c.action, c.val)
			if fk.OnDelete != c.wantDelete || fk.OnUpdate != c.wantUpdate {
				t.Errorf("setFKAction(%q,%q) -> OnDelete=%q OnUpdate=%q, want %q/%q",
					c.action, c.val, fk.OnDelete, fk.OnUpdate, c.wantDelete, c.wantUpdate)
			}
		})
	}
}
