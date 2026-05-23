//ff:type feature=validate type=model topic=manifest-infra
//ff:what TestValidateBuiltinBackendCase — table-driven 테스트 케이스 구조체

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

type TestValidateBuiltinBackendCase struct {
	name      string
	fs        *yongol.Fullstack
	spec      backendSpec
	wantCount int
}
