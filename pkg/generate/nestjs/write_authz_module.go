//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeAuthzModule — AuthzModule + AuthzService 파일 기록

package nestjs

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/authz"
)

// writeAuthzModule writes the AuthzModule and AuthzService files for DI.
func writeAuthzModule(srcDir string) error {
	authzDir := filepath.Join(srcDir, "authz")
	if err := os.MkdirAll(authzDir, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(authzDir, "authz.service.ts"), []byte(authz.RenderAuthzService()), 0o644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(authzDir, "authz.module.ts"), []byte(authz.RenderAuthzModule()), 0o644)
}
