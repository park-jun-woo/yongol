//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-statemachine
//ff:what groupByFile — entries 를 파일 경로로 bin 한다 (hurl-statemachine 용)

package hurl_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/hurl"

// groupByFile bins entries by source file so each file's sequence is
// reasoned about independently.
func groupByFile(entries []hurl.HurlEntry) map[string][]hurl.HurlEntry {
	out := map[string][]hurl.HurlEntry{}
	for _, e := range entries {
		out[e.File] = append(out[e.File], e)
	}
	return out
}
