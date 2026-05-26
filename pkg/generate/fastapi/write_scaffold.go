//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeScaffold — FastAPI 프로젝트 scaffold 파일 (pyproject.toml, requirements.txt, .env.example) 일괄 기록

package fastapi

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/fastapi/scaffold"
)

// writeScaffold writes pyproject.toml, requirements.txt, and .env.example.
func writeScaffold(backendDir, projectID string) error {
	if err := os.MkdirAll(backendDir, 0o755); err != nil {
		return err
	}
	pyproject, err := scaffold.RenderPyproject(projectID)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(backendDir, "pyproject.toml"), []byte(pyproject), 0o644); err != nil {
		return err
	}
	requirements, err := scaffold.RenderRequirements()
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(backendDir, "requirements.txt"), []byte(requirements), 0o644); err != nil {
		return err
	}
	envExample, err := scaffold.RenderEnvExample()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(backendDir, ".env.example"), []byte(envExample), 0o644)
}
