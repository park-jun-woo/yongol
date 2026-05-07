//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what s71ExtractPrefix — 값 표현에서 변수 prefix 추출 (리터럴·빈 문자열 skip)

package ssac

import (
	"strings"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func s71ExtractPrefix(val string) string {
	if val == "" {
		return ""
	}
	if strings.HasPrefix(val, `"`) || parsessac.IsLiteral(val) {
		return ""
	}
	if dot := strings.IndexByte(val, '.'); dot > 0 {
		return val[:dot]
	}
	return val
}
