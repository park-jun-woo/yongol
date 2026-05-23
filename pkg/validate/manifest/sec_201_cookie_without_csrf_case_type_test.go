//ff:type feature=validate type=model topic=manifest-auth
//ff:what TestSec201CookieWithoutCsrfCase — table-driven 테스트 케이스 구조체

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestSec201CookieWithoutCsrfCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
