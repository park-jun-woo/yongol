//ff:type feature=validate type=model topic=hurl-manifest
//ff:what TestProcessAuthEntryCase — table-driven 테스트 케이스 구조체

package hurl_manifest

import "github.com/park-jun-woo/yongol/pkg/parser/hurl"

type TestProcessAuthEntryCase struct {
	name          string
	entry         hurl.HurlEntry
	authIssued    bool
	wantAuth      bool
	wantDiagCount int
}
