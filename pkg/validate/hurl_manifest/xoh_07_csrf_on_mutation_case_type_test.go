//ff:type feature=validate type=model topic=hurl-manifest
//ff:what TestXoh07CSRFOnMutationCase — table-driven 테스트 케이스 구조체

package hurl_manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestXoh07CSRFOnMutationCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
	// wantMsgContains, when non-empty, must be a substring of every
	// diagnostic Message (e.g. the manifest-resolved CSRF header name).
	wantMsgContains string
}
