//ff:type feature=validate type=model topic=funcspec-structural
//ff:what TestRunCase — table-driven 테스트 케이스 구조체

package funcspec

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestRunCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
	wantCodes []string // expected diagnostic code substrings
}
