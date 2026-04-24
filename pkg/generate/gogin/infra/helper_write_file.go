//ff:func feature=gen-gogin type=util control=sequence
//ff:what writeAdapterFile — adapter 코드를 `<artifacts>/backend/internal/infra/<pkg>/postgres.go` 로 기록

package infra

import (
	"os"
	"path/filepath"
)

// writeAdapterFile creates the target directory (if needed) and writes the
// supplied Go source as `postgres.go`. Every wrapper emitter funnels its
// output through this helper so the artifact layout stays uniform.
func writeAdapterFile(artifactsDir, pkg string, content []byte) error {
	dir := filepath.Join(artifactsDir, "backend", "internal", "infra", pkg)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "postgres.go"), content, 0o644)
}
