//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what generateServerHelpers — Server 의 포인터/deref 유틸을 1파일 1func 로 분리 emit

package ssac

import (
	"os"
	"path/filepath"
)

// generateServerHelpers writes one file per pointer/deref helper used by
// service methods: ptr_of.go, deref_int.go, deref_str.go,
// deref_int64.go, deref_bool.go, deref_enum.go. Each file carries a
// complete //ff:func + //ff:what annotation so filefunc A1/A3/F1 pass
// uniformly across the service package.
func generateServerHelpers(artifactsDir string) error {
	dir := filepath.Join(artifactsDir, "backend", "internal", "service")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for fileName, body := range serverHelperSources() {
		if err := os.WriteFile(filepath.Join(dir, fileName), []byte(body), 0o644); err != nil {
			return err
		}
	}
	return nil
}
