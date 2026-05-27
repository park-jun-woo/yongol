//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what writeFuncStubs — 외부 패키지 stub Python 모듈 파일 기록

package fastapi

import (
	"fmt"
	"os"
	"path/filepath"

	funcstub "github.com/park-jun-woo/yongol/pkg/generate/fastapi/func"
)

// writeFuncStubs writes stub Python modules for each external package.
func writeFuncStubs(appDir string, pkgs []externalPackage) error {
	servicesDir := filepath.Join(appDir, "services")
	if err := os.MkdirAll(servicesDir, 0o755); err != nil {
		return fmt.Errorf("mkdir services: %w", err)
	}
	for _, pkg := range pkgs {
		content := funcstub.RenderFuncStub(pkg.Name, pkg.Methods)
		if err := os.WriteFile(filepath.Join(servicesDir, pkg.Name+".py"), []byte(content), 0o644); err != nil {
			return fmt.Errorf("write %s stub: %w", pkg.Name, err)
		}
	}
	return nil
}
