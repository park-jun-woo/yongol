//ff:func feature=gen-gogin type=util control=sequence
//ff:what generateOpenAPIGoGin — oapi-codegen -generate types,strict-server,gin 실행
package gogin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// generateOpenAPIGoGin invokes oapi-codegen with -generate types,strict-server,gin
// reading the OpenAPI spec at oapiPath and writing server.gen.go into outDir
// under the package pkgName. strict-server mode generates StrictServerInterface
// (no gin.Context in method signatures) + typed Request/Response objects + gin
// adapter. The SSaC codegen (pkg/generate/gogin/ssac/) implements
// StrictServerInterface with pure business logic — HTTP parsing/response is
// handled by the generated adapter. Single-site passes (specsDir/api/openapi.yaml,
// backend/internal/api, "api"); multi-domain passes per-domain values.
func generateOpenAPIGoGin(oapiPath, outDir, pkgName string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", outDir, err)
	}
	outFile := filepath.Join(outDir, "server.gen.go")

	cmd := exec.Command("oapi-codegen",
		"-package", pkgName,
		"-generate", "types,strict-server,gin",
		"-o", outFile,
		oapiPath,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w — %s", err, string(out))
	}
	return nil
}
