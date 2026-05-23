//ff:type feature=validate type=model topic=hurl-openapi
//ff:what TestXoh01URLMethodCase — table-driven 테스트 케이스 구조체

package hurl_openapi

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestXoh01URLMethodCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
