//ff:func feature=chain type=util control=iteration dimension=2
//ff:what traceOpenAPI finds the OpenAPI path/method for an operationId.
package chain

import (
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func traceOpenAPI(doc *openapi3.T, opID string, specsDir string) *Link {
	if doc.Paths == nil {
		slog.Debug("chain.traceOpenAPI: openapi doc.Paths nil", "operationId", opID)
		return nil
	}
	for path, pi := range doc.Paths.Map() {
		for method, op := range pi.Operations() {
			if op.OperationID == opID {
				line := grepLine(filepath.Join(specsDir, "api", "openapi.yaml"), "operationId: "+opID)
				return &Link{
					Kind:    "OpenAPI",
					File:    "api/openapi.yaml",
					Line:    line,
					Summary: strings.ToUpper(method) + " " + path,
				}
			}
		}
	}
	slog.Debug("chain.traceOpenAPI: operationId not matched in any path/method", "operationId", opID)
	return nil
}
