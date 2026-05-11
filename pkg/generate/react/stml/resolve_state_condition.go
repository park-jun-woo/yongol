//ff:func feature=stml-gen type=util control=selection
//ff:what StateBind의 Condition 문자열을 JSX 조건 표현식으로 변환한다
package stml

import (
	"fmt"
	"strings"
)

func resolveStateCondition(condition, dataVar string) string {
	switch {
	case strings.HasSuffix(condition, ".empty"):
		field := strings.TrimSuffix(condition, ".empty")
		return fmt.Sprintf("%s.%s?.length === 0", dataVar, field)
	case strings.HasSuffix(condition, ".loading"):
		return dataVar + "Loading"
	case strings.HasSuffix(condition, ".error"):
		return dataVar + "Error"
	default:
		return fmt.Sprintf("%s.%s", dataVar, condition)
	}
}
