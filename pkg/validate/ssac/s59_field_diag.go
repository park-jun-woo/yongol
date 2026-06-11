//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what s59FieldDiag — S-59 진단 생성: schema에 없는 필드를 case-insensitive 근사 일치로 did-you-mean 제시

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// s59FieldDiag builds the S-59 diagnostic for a missing variable field. When a
// field of the same name exists under different initialism casing (e.g. the
// SSaC text spells "QueueExportRepoURL" but the sqlc-generated struct field is
// "QueueExportRepoUrl"), it points the author at the canonical spelling instead
// of a bare "does not exist" (BUG-123). Near-match is detected via
// strings.EqualFold over the registered field list.
func s59FieldDiag(file string, line int, source, field string, schema []string) diagnostic.Diagnostic {
	d := diagnostic.Diagnostic{
		File:    file,
		Line:    line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[S-59] %s.%s: field %q does not exist on variable type", source, field, field),
		Advice:  fmt.Sprintf("Correct the field name to match an actual field of variable %s", source),
	}
	if suggestion, ok := findFieldFold(schema, field); ok {
		d.Message = fmt.Sprintf("[S-59] %s.%s: field %q does not exist — did you mean %q? (SSaC field names follow the sqlc-generated struct, e.g. Url not URL)", source, field, field, suggestion)
		d.Advice = fmt.Sprintf("Rename %s.%s to %s.%s", source, field, source, suggestion)
	}
	return d
}
