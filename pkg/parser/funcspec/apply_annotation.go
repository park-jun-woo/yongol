//ff:func feature=funcspec type=parser control=selection
//ff:what 코멘트 행에서 @func, @error, @description 값을 FuncSpec에 적용한다
package funcspec

import (
	"strconv"
	"strings"
)

func applyAnnotation(line string, lineNum int, spec *FuncSpec) {
	switch {
	case strings.HasPrefix(line, "@func "):
		spec.Name = strings.TrimSpace(strings.TrimPrefix(line, "@func "))
		spec.Line = lineNum
	case strings.HasPrefix(line, "@error "):
		if code, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "@error "))); err == nil {
			spec.ErrStatus = code
		}
	case strings.HasPrefix(line, "@description "):
		spec.Description = strings.TrimSpace(strings.TrimPrefix(line, "@description "))
	}
}
