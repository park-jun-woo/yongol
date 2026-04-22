//ff:func feature=validate-contract type=util control=sequence
//ff:what prv01MissingOpIDDiag — operationId 가 SSOT 에서 사라진 preserved 파일의 Diagnostic 생성

package contract

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// prv01MissingOpIDDiag formats the [PRV-01] diagnostic used when the
// preserved file's operationId (derived from its filename) is no
// longer declared in the OpenAPI SSOT. funcName may differ from opID
// when the user renamed the function while keeping the original
// filename.
func prv01MissingOpIDDiag(path, funcName, opID string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    path,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[PRV-01] %s — preserved, but SSOT has no operationId %q.", funcName, opID),
		Advice: strings.Join([]string{
			"(a) restore the operationId in specs/api/openapi.yaml",
			"(b) rename the preserved file to the new operationId (snake_case)",
			"(c) release preserve by deleting the file",
		}, "\n"),
	}
}
