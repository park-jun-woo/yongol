//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeFuncStubs — 외부 패키지 stub service+module 파일 기록

package nestjs

import (
	"fmt"
	"os"
	"path/filepath"

	funcstub "github.com/park-jun-woo/yongol/pkg/generate/nestjs/func"
)

// writeFuncStubs writes stub service and module files for each external package.
func writeFuncStubs(srcDir string, pkgs []externalPackage) error {
	for _, pkg := range pkgs {
		pkgDir := filepath.Join(srcDir, pkg.Name)
		if err := os.MkdirAll(pkgDir, 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", pkg.Name, err)
		}
		svcContent := funcstub.RenderFuncService(pkg.Name, pkg.Methods)
		if err := os.WriteFile(filepath.Join(pkgDir, pkg.Name+".service.ts"), []byte(svcContent), 0o644); err != nil {
			return fmt.Errorf("write %s service: %w", pkg.Name, err)
		}
		modContent := funcstub.RenderFuncModule(pkg.Name, pkg.Methods)
		if err := os.WriteFile(filepath.Join(pkgDir, pkg.Name+".module.ts"), []byte(modContent), 0o644); err != nil {
			return fmt.Errorf("write %s module: %w", pkg.Name, err)
		}
	}
	return nil
}
