//ff:func feature=validate type=util control=sequence topic=manifest-infra
//ff:what xdn04TypeMismatchDiag — XDN-04 단일 claim 타입 불일치 진단 메시지 생성

package manifest_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// xdn04TypeMismatchDiag formats the single-claim type-mismatch diagnostic
// emitted by XDN-04. Extracted so the parent rule's loop body stays under
// the Q4 pure-line budget.
func xdn04TypeMismatchDiag(field, userTable, claimGoType, ddlGoType string, def pmanifest.ClaimDef) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  "manifest.yaml",
		Line:  def.SourceLine,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDN-04] manifest auth.claims.%s declares Go type %q, but column %s.%s is %q in DDL",
			field, claimGoType, userTable, def.Key, ddlGoType,
		),
		Advice: fmt.Sprintf(
			"Change the claim mapping to `%s: %s:%s`, or change the DDL column type so the Go mapping matches.",
			field, def.Key, ddlGoType,
		),
	}
}
