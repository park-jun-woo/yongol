//ff:type feature=validate type=model topic=funcspec-structural
//ff:what TestCollectForbiddenImportsCase — table-driven 테스트 케이스 구조체

package funcspec

type TestCollectForbiddenImportsCase struct {
	name      string
	pkg       string
	funcName  string
	line      int
	imports   []string
	wantCount int
}
