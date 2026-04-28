//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what checkFileAuth — 한 hurl 파일의 entries 를 순회하며 인증 선행 WARNING 생성

package hurl_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// checkFileAuth walks a single file's entries and emits a WARNING for
// every protected call that precedes any auth step and lacks its own
// auth header.
func checkFileAuth(entries []hurl.HurlEntry) []diagnostic.Diagnostic {
	authIssued := false
	var diags []diagnostic.Diagnostic
	for _, e := range entries {
		authIssued = processAuthEntry(e, authIssued, &diags)
	}
	return diags
}
