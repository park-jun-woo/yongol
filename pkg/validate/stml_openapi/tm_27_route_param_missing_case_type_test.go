//ff:type feature=validate type=model topic=stml-openapi
//ff:what TestTM27RouteParamMissingCase — table-driven 테스트 케이스 구조체

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

type TestTM27RouteParamMissingCase struct {
	name      string
	page      stml.PageSpec
	wantCount int
}
