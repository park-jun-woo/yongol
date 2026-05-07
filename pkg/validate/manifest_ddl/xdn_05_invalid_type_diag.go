//ff:func feature=validate type=util control=sequence topic=manifest-infra
//ff:what xdn05InvalidTypeDiag — XDN-05 허용되지 않는 타입 선언 진단 메시지 생성

package manifest_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func xdn05InvalidTypeDiag(field string, def pmanifest.ClaimDef) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  "manifest.yaml",
		Line:  def.SourceLine,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDN-05] manifest auth.claims.%s declares unknown type %q",
			field, def.GoType,
		),
		Advice: "Allowed claim types: string, int64, int32, bool, uuid.",
	}
}
