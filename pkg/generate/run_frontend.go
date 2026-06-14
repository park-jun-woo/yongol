//ff:func feature=generate type=util control=selection
//ff:what runFrontend — FrontendType에 따라 React 스캐폴드 생성 + STML 페이지 생성 + 사용자 컴포넌트 복사 + tsc 게이트
package generate

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/react"
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
		// Final gate: type-check the generated frontend so "generate success =
		// buildable" holds for the frontend, symmetric to backend `go build`
		// (BUG-137 Phase041). Skips gracefully when the toolchain is unresolved.
		if err := react.RunTscCheck(artifactsDir); err != nil {
			return fmt.Errorf("tsc gate: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown frontend %q", frontend)
	}
}
