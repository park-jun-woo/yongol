//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what groupByFile — entries 를 파일 경로로 bin 한다

package hurl_manifest

import "github.com/park-jun-woo/yongol/pkg/parser/hurl"

// groupByFile bins entries by file so each hurl file is reasoned about
// independently. Hurl does not share captures across files, so auth
// state does not cross this boundary either.
func groupByFile(entries []hurl.HurlEntry) map[string][]hurl.HurlEntry {
	out := map[string][]hurl.HurlEntry{}
	for _, e := range entries {
		out[e.File] = append(out[e.File], e)
	}
	return out
}
