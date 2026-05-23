//ff:type feature=validate type=model topic=manifest-security-headers
//ff:what TestSec302HSTSShortCase — table-driven 테스트 케이스 구조체

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestSec302HSTSShortCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
