//ff:func feature=validate type=util control=sequence topic=manifest-infra
//ff:what xdn06TypeMismatchDiag — XDN-06 claim 타입 ↔ DDL 타입 불일치 진단 메시지 생성

package manifest_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func xdn06TypeMismatchDiag(field, userTable string, def pmanifest.ClaimDef, col ddl.Column) diagnostic.Diagnostic {
	suggested := suggestClaimType(col)
	return diagnostic.Diagnostic{
		File:  "manifest.yaml",
		Line:  def.SourceLine,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDN-06] manifest auth.claims.%s declares type %q, but column %s.%s is %q in DDL",
			field, def.GoType, userTable, def.Key, col.RawType,
		),
		Advice: fmt.Sprintf(
			"Change the claim mapping to `%s: %s:%s`, or change the DDL column type.",
			field, def.Key, suggested,
		),
	}
}
