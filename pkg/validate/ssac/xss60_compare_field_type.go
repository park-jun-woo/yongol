//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what xss60CompareFieldType — subscriber 필드 타입을 publish 필드 타입과 비교하여 불일치 시 진단 반환

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xss60CompareFieldType compares a single subscriber field against the publish
// type map. Returns (diag, true) on mismatch.
func xss60CompareFieldType(fn parsessac.ServiceFunc, field parsessac.StructField, pubFields map[string]string) (diagnostic.Diagnostic, bool) {
	pubType, ok := pubFields[field.Name]
	if !ok || pubType == "" || pubType == field.Type {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:  fn.FileName,
		Line:  fn.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelWarning,
		Message: fmt.Sprintf(
			"[XSS-60] topic %q field %q — publisher sends %s but subscriber expects %s",
			fn.Subscribe.Topic, field.Name, pubType, field.Type,
		),
		Advice: "Align subscriber field type with publisher payload type",
	}, true
}
