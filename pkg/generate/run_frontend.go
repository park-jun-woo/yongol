//ff:func feature=generate type=util control=selection
//ff:what runFrontend — FrontendType에 따라 React 스캐폴드 생성 + 사용자 컴포넌트 복사
package generate

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/generate/react"
)

// runFrontend executes the frontend generator pipeline for the given
// target: react.Generate emits api.ts + build configs + App/main scaffolds,
// then user-authored components are mirrored into the output tree.
func runFrontend(fs *yongol.Fullstack, artifactsDir string, frontend FrontendType) error {
	switch frontend {
	case React:
		if err := react.Generate(fs, artifactsDir); err != nil {
			return err
		}
		if err := copyFrontendComponents(fs.SpecsDir, artifactsDir); err != nil {
			return fmt.Errorf("components: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown frontend %q", frontend)
	}
}
