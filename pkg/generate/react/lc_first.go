//ff:func feature=gen-react type=util control=sequence
//ff:what camelCase 변환 유틸리티 함수

package react

import "github.com/ettle/strcase"

func lcFirst(s string) string {
	return strcase.ToGoCamel(s)
}
