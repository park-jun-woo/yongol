//ff:func feature=util type=test control=iteration dimension=1 topic=string-convert
//ff:what KebabToCamel 회귀 테이블 테스트

package caseconv

import "testing"

func TestKebabToCamel(t *testing.T) {
	cases := []struct{ in, want string }{
		{"project-id", "projectId"},
		{"ReservationID", "ReservationID"},
		{"room-id", "roomId"},
		{"a-b-c", "aBC"},
		{"data-fetch", "dataFetch"},
		{"nodash", "nodash"},
		{"", ""},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := KebabToCamel(c.in); got != c.want {
				t.Errorf("KebabToCamel(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
