//ff:type feature=validate type=model topic=manifest-structural
//ff:what TestC08RateLimitLoginRequiredCase — table-driven 테스트 케이스 구조체

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestC08RateLimitLoginRequiredCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
