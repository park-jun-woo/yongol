//ff:func feature=chain type=util control=iteration dimension=2
//ff:what traceHurl finds .hurl files referencing the given endpoint.
package chain

import (
	"log/slog"
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"
)

func traceHurl(opID string, doc *openapi3.T, testsDir string, specsDir string) []Link {
	if doc == nil || doc.Paths == nil {
		slog.Debug("chain.traceHurl: openapi doc or paths nil", "operationId", opID)
		return nil
	}

	var endpointPath string
	for path, pi := range doc.Paths.Map() {
		for _, op := range pi.Operations() {
			if op.OperationID == opID {
				endpointPath = path
				break
			}
		}
		if endpointPath != "" {
			break
		}
	}
	if endpointPath == "" {
		slog.Debug("chain.traceHurl: operationId not found in openapi paths", "operationId", opID)
		return nil
	}

	var links []Link
	hurlFiles, _ := filepath.Glob(filepath.Join(testsDir, "*.hurl"))
	for _, f := range hurlFiles {
		line := grepLine(f, endpointPath)
		if line > 0 {
			relPath, _ := filepath.Rel(specsDir, f)
			links = append(links, Link{
				Kind:    "Hurl",
				File:    relPath,
				Line:    line,
				Summary: "scenario: " + filepath.Base(f),
			})
		}
	}
	return links
}
