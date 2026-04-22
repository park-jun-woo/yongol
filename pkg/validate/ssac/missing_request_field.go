//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what missingRequestField — S-60 ERROR Diagnostic 생성 헬퍼

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// missingRequestField builds the S-60 ERROR Diagnostic for a request field
// that is referenced by SSaC but absent from the OpenAPI request schema.
func missingRequestField(fn ssac.ServiceFunc, seq ssac.Sequence, field string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    fn.FileName,
		Line:    seq.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[S-60] request.%s not in OpenAPI request schema", field),
		Advice:  fmt.Sprintf("OpenAPI 스키마에 field %q 를 case-exact 로 추가하거나 SSaC 에서 정확한 snake_case 필드명으로 변경하세요", field),
	}
}
