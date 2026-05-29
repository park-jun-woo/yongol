//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what collectFuncs — 활성 블록의 Funcs 를 순서대로 수집

package boot

import (
	"strings"
	"testing"
)

func TestCollectFuncs(t *testing.T) {
	cases := []struct {
		name   string
		blocks []MainBlock
		want   []string
	}{
		{"empty", nil, nil},
		{"no funcs", []MainBlock{{Lines: []string{"x"}}}, nil},
		{
			"preserves order and duplicates",
			[]MainBlock{
				{Funcs: []string{"funcA", "funcB"}},
				{Funcs: []string{"funcA"}},
			},
			[]string{"funcA", "funcB", "funcA"},
		},
	}
	for _, c := range cases {
		got := collectFuncs(c.blocks)
		if strings.Join(got, "|") != strings.Join(c.want, "|") {
			t.Errorf("%s: collectFuncs = %v, want %v", c.name, got, c.want)
		}
	}
}
