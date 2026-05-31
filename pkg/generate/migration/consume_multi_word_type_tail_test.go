//ff:func feature=migration type=test control=iteration dimension=1
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestConsumeMultiWordTypeTail(t *testing.T) {
	cases := []struct {
		name      string
		toks      []string
		i         int
		startPart []string
		wantParts []string
		wantIdx   int
	}{
		{"varying", []string{"character", "varying"}, 1, []string{"character"}, []string{"character", "varying"}, 2},
		{"precision", []string{"double", "precision"}, 1, []string{"double"}, []string{"double", "precision"}, 2},
		{"with time zone", []string{"timestamp", "with", "time", "zone"}, 1, []string{"timestamp"}, []string{"timestamp", "with", "time", "zone"}, 4},
		{"without time zone", []string{"timestamp", "without", "time", "zone"}, 1, []string{"timestamp"}, []string{"timestamp", "without", "time", "zone"}, 4},
		{"no tail match", []string{"timestamp", "NOT"}, 1, []string{"timestamp"}, []string{"timestamp"}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertConsumeMultiWordTypeTail(t, c.toks, c.i, c.startPart, c.wantParts, c.wantIdx)
		})
	}
}
