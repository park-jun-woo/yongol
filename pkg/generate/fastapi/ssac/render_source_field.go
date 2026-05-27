//ff:func feature=gen-fastapi type=util control=selection
//ff:what renderSourceField — source 유형별 source.field → Python 표현식 변환

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
			return fmt.Sprintf("params[\"%s\"]", snakeCase(strings.TrimPrefix(field, "Params.")))
		}
		return fmt.Sprintf("params[\"%s\"]", snakeCase(field))
	case "currentUser":
		return fmt.Sprintf("user[\"%s\"]", snakeCase(field))
	default:
		return fmt.Sprintf("%s[\"%s\"]", source, snakeCase(field))
	}
}
