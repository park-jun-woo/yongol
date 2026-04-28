//ff:func feature=validate type=util control=sequence topic=manifest-infra
//ff:what xdn03MissingColumnDiag — XDN-03 단일 claim 진단 메시지 생성

package manifest_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// xdn03MissingColumnDiag formats the single-claim "column not found"
// diagnostic emitted by XDN-03. Extracted from xdn03ClaimColumnExists so
// the loop body fits the Q4 pure-line budget (filefunc).
func xdn03MissingColumnDiag(field, userTable string, def pmanifest.ClaimDef) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  "manifest.yaml",
		Line:  def.SourceLine,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDN-03] manifest auth.claims.%s: %s, but column %q not found in table %q",
			field, def.Key, def.Key, userTable,
		),
		Advice: "Add the column to " + userTable + ", or remove the claim from " +
			"backend.auth.claims.",
	}
}
