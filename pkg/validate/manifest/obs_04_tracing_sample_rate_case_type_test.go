//ff:type feature=validate type=model topic=manifest-observability
//ff:what TestObs04TracingSampleRateCase — table-driven 테스트 케이스 구조체

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestObs04TracingSampleRateCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
