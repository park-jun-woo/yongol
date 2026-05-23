//ff:type feature=validate type=model topic=manifest-cors
//ff:what TestCors01WildcardCredentialsCase — table-driven 테스트 케이스 구조체

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestCors01WildcardCredentialsCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
