//ff:func feature=generate type=util control=selection
//ff:what runFrontend — FrontendType에 따라 React 스캐폴드 생성 + STML 페이지 생성 + 사용자 컴포넌트 복사
package generate

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/react"
	stmlgen "github.com/park-jun-woo/yongol/pkg/generate/react/stml"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runFrontend executes the frontend generator pipeline for the given
// target: react.Generate emits api.ts + build configs + App/main scaffolds,
// STML codegen produces page TSX files from STMLPages when present,
// then user-authored components are mirrored into the output tree.
func runFrontend(fs *yongol.Fullstack, artifactsDir string, frontend FrontendType) error {
	switch frontend {
	case React:
		if err := react.Generate(fs, artifactsDir); err != nil {
			return err
		}
		if err := runSTMLCodegen(fs, artifactsDir); err != nil {
			return fmt.Errorf("stml: %w", err)
		}
		if err := copyFrontendComponents(fs.SpecsDir, artifactsDir); err != nil {
			return fmt.Errorf("components: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown frontend %q", frontend)
	}
}

// runSTMLCodegen generates page TSX files from STML specs when STMLPages
// are present. Skips silently if no STML pages were parsed.
func runSTMLCodegen(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || len(fs.STMLPages) == 0 {
		return nil
	}
	pagesDir := filepath.Join(artifactsDir, "frontend", "src", "pages")
	opt := stmlgen.GenerateOptions{
		RequestConstraints:      fs.RequestConstraints,
		ResponseArrayItemFields: oapiparser.ExtractResponseArrayItemFields(fs.OpenAPIDoc),
	}
	_, err := stmlgen.Generate(fs.STMLPages, fs.SpecsDir, pagesDir, opt)
	return err
}
