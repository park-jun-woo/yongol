//ff:func feature=external type=util control=sequence
//ff:what 문자열을 Go camelCase로 변환한다
package external

import "github.com/ettle/strcase"

func toCamelCase(s string) string {
	return strcase.ToGoCamel(s)
}
