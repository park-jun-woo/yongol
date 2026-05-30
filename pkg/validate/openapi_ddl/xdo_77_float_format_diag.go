//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what xdo77FloatFormatDiag — float DDL 컬럼에 OpenAPI format: double 누락 시 진단 메시지 생성

package openapi_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// xdo77FloatFormatDiag builds the float-specific diagnostic for XDO-77 when the
// DDL column maps to Go float64 but the OpenAPI property lacks format: double.
// yongol's GoTypeOf projects every float column (NUMERIC/DECIMAL/REAL/FLOAT/
// FLOAT4/FLOAT8) to float64, so a formatless OpenAPI `number` (oapi-codegen
// float32) never matches the sqlc/response float64 and breaks the generated build.
func xdo77FloatFormatDiag(tableName, colName string, line int) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  "api/openapi.yaml",
		Line:  line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDO-77] DDL column %s.%s maps to float64 but OpenAPI field has type: number without format: double",
			tableName, colName,
		),
		Advice: "Add format: double to the OpenAPI property (yongol maps all float columns to float64)",
	}
}
