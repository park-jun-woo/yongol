//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeScaffold — NestJS 프로젝트 scaffold 파일 (package.json, tsconfig, nest-cli) 일괄 기록

package nestjs

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/scaffold"
)

// writeScaffold writes package.json, tsconfig.json, and nest-cli.json.
func writeScaffold(backendDir, projectID string) error {
	if err := os.MkdirAll(backendDir, 0o755); err != nil {
		return err
	}
	pkgJSON, err := scaffold.RenderPackageJSON(projectID)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(backendDir, "package.json"), []byte(pkgJSON), 0o644); err != nil {
		return err
	}
	tsConfig, err := scaffold.RenderTSConfig()
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(backendDir, "tsconfig.json"), []byte(tsConfig), 0o644); err != nil {
		return err
	}
	nestCLI, err := scaffold.RenderNestCLI()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(backendDir, "nest-cli.json"), []byte(nestCLI), 0o644)
}
