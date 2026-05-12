//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what 문자열 슬라이스를 작은따옴표로 감싸고 | 로 연결한다

package react

import (
	"fmt"
	"strings"
)

// quotedUnion formats a list of keys as a TypeScript union type string.
// e.g. ["a", "b"] → "'a' | 'b'"
func quotedUnion(keys []string) string {
	quoted := make([]string, len(keys))
	for i, k := range keys {
		quoted[i] = fmt.Sprintf("'%s'", k)
	}
	return strings.Join(quoted, " | ")
}
