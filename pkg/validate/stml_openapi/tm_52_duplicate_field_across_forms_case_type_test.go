//ff:type feature=validate type=model topic=stml-openapi
//ff:what tm52Case — table-driven 테스트 케이스 구조체

package stml_openapi

type tm52Case struct {
	name      string
	html      string
	wantCount int
	wantField string
}
