//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what emitIfClass — className 이 비어있지 않으면 TM-10 진단 생성

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// emitIfClass returns a TM-10 diagnostic if className is non-empty.
func emitIfClass(file, elemType, elemID, className string) []diagnostic.Diagnostic {
	if className == "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-10] class attribute %q on %s %q is prohibited in STML; use <!-- @override class=\"...\" --> comment instead", className, elemType, elemID),
		Advice:  "Remove the class attribute from the STML element. To override DESIGN.md styling, place <!-- @override class=\"...\" --> as a comment before the element",
	}}
}
