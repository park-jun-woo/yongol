//ff:type feature=validate type=model topic=hurl-structural
//ff:what TestH02EmptyTestsDirCase — table-driven 테스트 케이스 구조체

package hurl

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestH02EmptyTestsDirCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
