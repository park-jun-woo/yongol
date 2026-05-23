//ff:type feature=validate type=model topic=manifest-security-headers
//ff:what TestSec301CspPermissiveCase — table-driven 테스트 케이스 구조체

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestSec301CspPermissiveCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
