//ff:type feature=validate type=model topic=funcspec-structural
//ff:what TestF01BuiltinOverrideCase — table-driven 테스트 케이스 구조체

package funcspec

import "github.com/park-jun-woo/yongol/pkg/parser/funcspec"

type TestF01BuiltinOverrideCase struct {
	name      string
	builtin   []funcspec.FuncSpec
	project   []funcspec.FuncSpec
	wantCount int
}
