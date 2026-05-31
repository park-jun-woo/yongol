//ff:func feature=migration type=test control=iteration dimension=1
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"reflect"
	"testing"
)

func TestCollectTypeTokens(t *testing.T) {
	cases := []struct {
		name     string
		in       []string
		wantType string
		wantRest []string
	}{
		{"empty", nil, "", nil},
		{"simple", []string{"INTEGER", "NOT", "NULL"}, "INTEGER", []string{"NOT", "NULL"}},
		{"varchar with len token", []string{"VARCHAR(255)", "NOT"}, "VARCHAR(255)", []string{"NOT"}},
		{"character varying", []string{"character", "varying", "NOT"}, "character varying", []string{"NOT"}},
		{"timestamp with time zone", []string{"timestamp", "with", "time", "zone"}, "timestamp with time zone", []string{}},
		{"timestamp without time zone", []string{"timestamp", "without", "time", "zone", "NOT"}, "timestamp without time zone", []string{"NOT"}},
		{"double precision", []string{"double", "precision"}, "double precision", []string{}},
		{"array", []string{"INTEGER", "[]"}, "INTEGER[]", []string{}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			gotType, gotRest := collectTypeTokens(c.in)
			if gotType != c.wantType {
				t.Errorf("type = %q, want %q", gotType, c.wantType)
			}
			if !reflect.DeepEqual(gotRest, c.wantRest) {
				t.Errorf("rest = %#v, want %#v", gotRest, c.wantRest)
			}
		})
	}
}
