//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what groupEntriesByFile — entries 를 파일 경로로 bin 한다 (hurl-openapi 용)

package hurl_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/hurl"

// groupEntriesByFile bins entries by their source file. Preserves
// declaration order within each bucket so downstream rules can do
// positional analysis.
func groupEntriesByFile(entries []hurl.HurlEntry) map[string][]hurl.HurlEntry {
	out := map[string][]hurl.HurlEntry{}
	for _, e := range entries {
		out[e.File] = append(out[e.File], e)
	}
	return out
}
