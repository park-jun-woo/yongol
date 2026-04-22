//ff:func feature=funcspec type=util control=sequence
//ff:what AST 필드에서 json 태그 값을 추출한다
package funcspec

import (
	"go/ast"
	"reflect"
	"strings"
)

func extractJSONTag(f *ast.Field) string {
	if f.Tag == nil {
		return ""
	}
	tag := reflect.StructTag(strings.Trim(f.Tag.Value, "`"))
	jn, ok := tag.Lookup("json")
	if !ok {
		return ""
	}
	jn = strings.Split(jn, ",")[0]
	if jn == "" || jn == "-" {
		return ""
	}
	return jn
}
