//ff:type feature=validate type=model topic=stml-openapi
//ff:what TestTM30ItemSourceCase — table-driven 테스트 케이스 구조체

package stml_openapi

type TestTM30ItemSourceCase struct {
	name      string
	html      string
	raif      map[string]map[string]map[string]bool
	wantCount int
}
