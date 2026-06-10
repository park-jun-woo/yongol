//ff:type feature=validate type=model topic=stml-openapi
//ff:what TestTM31LinkTargetCase — table-driven 테스트 케이스 구조체

package stml_openapi

type TestTM31LinkTargetCase struct {
	name      string
	html      string
	pageNames []string
	wantCount int
}
