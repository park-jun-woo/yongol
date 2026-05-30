//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what matchMultiTokenHead — 선두 토큰 일치 시 소비 수, 불일치/짧음 0

package ddl

import "testing"

func TestMatchMultiTokenHead(t *testing.T) {
	cases := []struct {
		name  string
		upper []string
		head  []string
		want  int
	}{
		{"exact match", []string{"DOUBLE", "PRECISION"}, []string{"DOUBLE", "PRECISION"}, 2},
		{"match with trailing param", []string{"CHARACTER", "VARYING(255),"}, []string{"CHARACTER", "VARYING"}, 2},
		{"upper too short", []string{"DOUBLE"}, []string{"DOUBLE", "PRECISION"}, 0},
		{"mismatch first", []string{"BIG", "INT"}, []string{"DOUBLE", "PRECISION"}, 0},
		{"single token head", []string{"BIGINT", "NOT", "NULL"}, []string{"BIGINT"}, 1},
		{"mismatch last after strip", []string{"CHARACTER", "FIXED"}, []string{"CHARACTER", "VARYING"}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := matchMultiTokenHead(c.upper, c.head); got != c.want {
				t.Errorf("matchMultiTokenHead(%v,%v) = %d, want %d", c.upper, c.head, got, c.want)
			}
		})
	}
}
