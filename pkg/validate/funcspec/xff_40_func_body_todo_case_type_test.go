//ff:type feature=validate type=model topic=funcspec-structural
//ff:what TestXff40FuncBodyTodoCase — table-driven 테스트 케이스 구조체

package funcspec

import "github.com/park-jun-woo/yongol/pkg/parser/funcspec"

type TestXff40FuncBodyTodoCase struct {
	name      string
	specs     []funcspec.FuncSpec
	wantCount int
}
