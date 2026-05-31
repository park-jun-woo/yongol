//ff:func feature=ground type=test control=iteration dimension=1 topic=ddl
//ff:what populateDDL apifield — DDL.apifield.<M>.<f> 등록 + Struct.<M>.<f> 키 토큰 동일성 고정 (BUG-099)
package ground

import (
	"strings"
	"testing"
)

func keyTokens(t *testing.T, types map[string]string, prefix string) map[string]bool {
	t.Helper()
	out := make(map[string]bool)
	for k := range types {
		if strings.HasPrefix(k, prefix) {
			out[strings.TrimPrefix(k, prefix)] = true
		}
	}
	return out
}
