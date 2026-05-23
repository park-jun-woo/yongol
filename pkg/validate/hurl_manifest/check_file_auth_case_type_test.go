//ff:type feature=validate type=model topic=hurl-manifest
//ff:what TestCheckFileAuthCase — table-driven 테스트 케이스 구조체

package hurl_manifest

import "github.com/park-jun-woo/yongol/pkg/parser/hurl"

type TestCheckFileAuthCase struct {
	name      string
	entries   []hurl.HurlEntry
	wantCount int
}
