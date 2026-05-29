//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what assembleBody — 활성 블록의 Lines 를 순서대로 연결

package boot

import (
	"strings"
	"testing"
)

func TestAssembleBody(t *testing.T) {
	cases := []struct {
		name   string
		blocks []MainBlock
		want   []string
	}{
		{"empty", nil, nil},
		{
			"single block",
			[]MainBlock{{Lines: []string{"a", "b"}}},
			[]string{"a", "b"},
		},
		{
			"two blocks separated by blank line",
			[]MainBlock{{Lines: []string{"a"}}, {Lines: []string{"b"}}},
			[]string{"a", "", "b"},
		},
		{
			"empty-lined block does not add separator",
			[]MainBlock{{Lines: []string{"a"}}, {Lines: nil}, {Lines: []string{"b"}}},
			[]string{"a", "", "b"},
		},
	}
	for _, c := range cases {
		got := assembleBody(c.blocks)
		if strings.Join(got, "|") != strings.Join(c.want, "|") {
			t.Errorf("%s: assembleBody = %v, want %v", c.name, got, c.want)
		}
	}
}
