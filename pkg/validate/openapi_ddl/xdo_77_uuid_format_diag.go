//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what xdo77UUIDFormatDiag — UUID DDL 컬럼에 OpenAPI format: uuid 누락 시 진단 메시지 생성

package openapi_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// xdo77UUIDFormatDiag builds the UUID-specific diagnostic for XDO-77 when the
// DDL column type is UUID but the OpenAPI property lacks format: uuid.
func xdo77UUIDFormatDiag(tableName, colName string, line int) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  "api/openapi.yaml",
		Line:  line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDO-77] DDL column %s.%s is UUID but OpenAPI field has type: string without format: uuid",
			tableName, colName,
		),
		Advice: "Add format: uuid to the OpenAPI property",
	}
}
