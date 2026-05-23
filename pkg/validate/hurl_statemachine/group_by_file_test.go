//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-statemachine
//ff:what groupByFile — entries를 파일별 분류 검증

package hurl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestGroupByFile(t *testing.T) {
	cases := []struct {
		name       string
		entries    []hurl.HurlEntry
		wantFiles  int
		wantCounts map[string]int
	}{
		{name: "nil", entries: nil, wantFiles: 0},
		{
			name:       "single_file",
			entries:    []hurl.HurlEntry{{File: "a.hurl"}, {File: "a.hurl"}},
			wantFiles:  1,
			wantCounts: map[string]int{"a.hurl": 2},
		},
		{
			name:       "multiple_files",
			entries:    []hurl.HurlEntry{{File: "a.hurl"}, {File: "b.hurl"}, {File: "a.hurl"}},
			wantFiles:  2,
			wantCounts: map[string]int{"a.hurl": 2, "b.hurl": 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runFileGroupCase(t, groupByFile(c.entries), c.wantFiles, c.wantCounts)
		})
	}
}
