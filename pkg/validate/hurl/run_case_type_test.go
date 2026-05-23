//ff:type feature=validate type=model topic=hurl-structural
//ff:what TestRunCase — table-driven 테스트 케이스 구조체

package hurl

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestRunCase struct {
	name      string
	setup     func(dir string)
	presences map[yongol.SSOTKind]yongol.SSOTPresence
	wantCodes []string
}
