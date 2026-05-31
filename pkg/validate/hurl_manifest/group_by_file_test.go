//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what groupByFile — entries를 파일별로 분류하는 로직 검증

package hurl_manifest

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
		{
			name:      "nil_entries",
			entries:   nil,
			wantFiles: 0,
		},
		{
			name: "single_file",
			entries: []hurl.HurlEntry{
				{File: "a.hurl", Method: "GET", Path: "/x"},
				{File: "a.hurl", Method: "POST", Path: "/y"},
			},
			wantFiles:  1,
			wantCounts: map[string]int{"a.hurl": 2},
		},
		{
			name: "multiple_files",
			entries: []hurl.HurlEntry{
				{File: "a.hurl", Method: "GET", Path: "/x"},
				{File: "b.hurl", Method: "POST", Path: "/y"},
				{File: "a.hurl", Method: "DELETE", Path: "/z"},
			},
			wantFiles:  2,
			wantCounts: map[string]int{"a.hurl": 2, "b.hurl": 1},
		},
		{
			name:      "empty_entries",
			entries:   []hurl.HurlEntry{},
			wantFiles: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runFileGroupCase(t, groupByFile(c.entries), c.wantFiles, c.wantCounts)
		})
	}
}
