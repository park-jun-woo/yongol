//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what goStringSlice — []string을 Go []string{...} 리터럴로 렌더링

package boot

import (
	"fmt"
	"strings"
)

// goStringSlice renders a []string as a Go literal. Empty slice becomes
// `[]string{}` for assignment safety.
func goStringSlice(ss []string) string {
	if len(ss) == 0 {
		return `[]string{}`
	}
	quoted := make([]string, len(ss))
	for i, s := range ss {
		quoted[i] = fmt.Sprintf("%q", s)
	}
	return "[]string{" + strings.Join(quoted, ", ") + "}"
}
