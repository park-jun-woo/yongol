//ff:func feature=validate type=util control=sequence topic=manifest-infra
//ff:what xdn05MissingTypeDiag — XDN-05 타입 미선언 진단 메시지 생성

package manifest_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func xdn05MissingTypeDiag(field string, def pmanifest.ClaimDef) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  "manifest.yaml",
		Line:  def.SourceLine,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDN-05] manifest auth.claims.%s value %q must declare a type (col:type format required)",
			field, def.Key,
		),
		Advice: fmt.Sprintf(
			"Change to `%s: %s:<type>` where type is one of: string, int64, int32, bool, uuid.",
			field, def.Key,
		),
	}
}
