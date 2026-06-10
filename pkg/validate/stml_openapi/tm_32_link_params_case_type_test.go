//ff:type feature=validate type=model topic=stml-openapi
//ff:what TestTM32LinkParamsCase — table-driven 테스트 케이스 구조체

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

type TestTM32LinkParamsCase struct {
	name      string
	html      string
	targets   []stml.PageSpec
	raif      map[string]map[string]map[string]bool
	wantCount int
	wantIn    string // substring expected in some diagnostic message (empty = skip)
}
