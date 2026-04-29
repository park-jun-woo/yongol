//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what buildXqs20Diag — XQS-20 진단 메시지 + advice 조립

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildXqs20Diag composes a XQS-20 ERROR diagnostic for a sequence whose
// declared return type does not match the sqlc query's RETURNING shape.
//
//   - declared   : the type the user wrote in SSaC (e.g. "User", "UserNewRow")
//   - expected   : the type yongol expects for the query's actual shape
//   - queryName  : the sqlc query name (e.g. "UserCreate")
//   - shape      : the classified shape ("full" / "partial")
//   - reason     : short reason fragment shown in the message
//
// Advice is direction-aware:
//   - declared = Model + actual partial → suggest switching to <QueryName>Row
//   - declared = Row   + actual full    → suggest switching to <Model>
func buildXqs20Diag(
	fn ssac.ServiceFunc,
	seq ssac.Sequence,
	declared, expected, queryName string,
	shape ReturningShape,
	reason string,
) diagnostic.Diagnostic {
	advice := buildXqs20Advice(seq.Type, declared, expected, shape)
	return diagnostic.Diagnostic{
		File:    fn.FileName,
		Line:    seq.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf(`[XQS-20] @%s declares return type %q but sqlc query %q returns %q (%s)`, seq.Type, declared, queryName, expected, reason),
		Advice:  advice,
	}
}
