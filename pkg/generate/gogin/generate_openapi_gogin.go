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
// and writes the output to artifactsDir/backend/internal/api/server.gen.go.
// strict-server mode generates StrictServerInterface (no gin.Context in method
// signatures) + typed Request/Response objects + gin adapter. The SSaC codegen
// (pkg/generate/gogin/ssac/) implements StrictServerInterface with pure
// business logic — HTTP parsing/response is handled by the generated adapter.
func generateOpenAPIGoGin(specsDir, artifactsDir string) error {
	apiPath := filepath.Join(specsDir, "api", "openapi.yaml")
	outDir := filepath.Join(artifactsDir, "backend", "internal", "api")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", outDir, err)
	}
	outFile := filepath.Join(outDir, "server.gen.go")

	cmd := exec.Command("oapi-codegen",
		"-package", "api",
		"-generate", "types,strict-server,gin",
		"-o", outFile,
		apiPath,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w — %s", err, string(out))
	}
	return nil
}
