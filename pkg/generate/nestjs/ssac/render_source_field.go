//ff:func feature=gen-nestjs type=util control=selection
//ff:what renderSourceField — source 유형별 source.field → TypeScript 표현식 변환

package ssac

import (
	"fmt"
	"strings"
)

// renderSourceField dispatches source.field rendering by source type.
func renderSourceField(source, field string) string {
	switch source {
	case "request":
		if strings.HasPrefix(field, "Params.") {
			return fmt.Sprintf("query.%s", toSnake(strings.TrimPrefix(field, "Params.")))
		}
		return fmt.Sprintf("params.%s", toSnake(field))
	case "currentUser":
		return fmt.Sprintf("user.%s", toSnake(field))
	default:
		return fmt.Sprintf("%s.%s", source, toSnake(field))
	}
}
