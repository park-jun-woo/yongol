//ff:func feature=generate type=util control=selection
//ff:what runFrontend — FrontendType에 따라 React 스캐폴드 생성 + STML 페이지 생성 + 사용자 컴포넌트 복사 + tsc 게이트
package generate

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/react"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runFrontend executes the frontend generator pipeline for the given
// target: react.Generate emits api.ts + build configs + App/main scaffolds,
// STML codegen produces page TSX files from STMLPages when present,
// then user-authored components are mirrored into the output tree and the
// generated frontend is type-checked (tsc gate).
//
// In domain mode (fs.IsDomained()) react.Generate already scaffolds one app
// per domain under <artifacts>/frontend/<name> (Phase009); the post-steps here
// (STML codegen + component copy + tsc gate) are likewise run once per domain
// against that domain directory (Phase010, Decision A/N). Single-site behavior
// is unchanged: one pipeline against <artifacts>/frontend.
func runFrontend(fs *yongol.Fullstack, artifactsDir string, frontend FrontendType) error {
	switch frontend {
	case None:
		return nil
	case React:
		if err := react.Generate(fs, artifactsDir); err != nil {
			return err
		}
		if fs.IsDomained() {
			return runDomainFrontendPipelines(fs, artifactsDir)
		}
		frontendDir := filepath.Join(artifactsDir, "frontend")
		if err := runSTMLCodegen(fs, fs.SpecsDir, frontendDir); err != nil {
			return fmt.Errorf("stml: %w", err)
		}
		if err := copyFrontendComponents(filepath.Join(fs.SpecsDir, "frontend"), filepath.Join(frontendDir, "src")); err != nil {
			return fmt.Errorf("components: %w", err)
		}
		// Final gate: type-check the generated frontend so "generate success =
		// buildable" holds for the frontend, symmetric to backend `go build`
		// (BUG-137 Phase041). Skips gracefully when the toolchain is unresolved.
		if err := react.RunTscCheck(frontendDir); err != nil {
			return fmt.Errorf("tsc gate: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown frontend %q", frontend)
	}
}
