//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what writeComponentsUI — DESIGN.md components 기반 또는 하드코딩 fallback 으로 src/components/ui/ 방출

package react

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// writeComponentsUI materializes UI component files under src/components/ui/.
// When a DesignSpec with components is provided, each component is rendered
// from its base/variants/sizes definition. Otherwise, the 10 hardcoded
// shadcn-like primitives are emitted as a fallback.
func writeComponentsUI(srcDir string, spec *design.DesignSpec) error {
	uiDir := filepath.Join(srcDir, "components", "ui")
	if err := os.MkdirAll(uiDir, 0o755); err != nil {
		return err
	}

	if spec != nil && len(spec.Components) > 0 {
		return writeDesignComponents(uiDir, spec.Components)
	}

	for name, source := range uiPrimitives() {
		if err := os.WriteFile(filepath.Join(uiDir, name), []byte(source), 0o644); err != nil {
			return err
		}
	}
	return nil
}
