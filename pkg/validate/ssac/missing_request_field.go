//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what missingRequestField — helper that generates an S-60 ERROR Diagnostic

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
		Advice:  fmt.Sprintf("Add field %q to the OpenAPI schema with the exact case, or change the SSaC reference to the exact snake_case field name", field),
	}
}
