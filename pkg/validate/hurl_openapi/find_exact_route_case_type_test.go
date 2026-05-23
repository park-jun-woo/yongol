//ff:type feature=validate type=model topic=hurl-openapi
//ff:what TestFindExactRouteCase — table-driven 테스트 케이스 구조체

package hurl_openapi

type TestFindExactRouteCase struct {
	name     string
	segs     []string
	method   string
	wantPath string
	wantNil  bool
}
