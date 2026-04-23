//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what writeComponentsUI — src/components/ui/ 아래 shadcn-like 프리미티브 10종 방출

package react

import (
	"os"
	"path/filepath"
)

// writeComponentsUI materializes 10 shadcn/ui-inspired primitives under
// src/components/ui/. Each file is self-contained and only depends on
// React + @/lib/utils (cn). Variants follow shadcn conventions so AI can
// iterate without extra lookup cost.
func writeComponentsUI(srcDir string) error {
	uiDir := filepath.Join(srcDir, "components", "ui")
	if err := os.MkdirAll(uiDir, 0o755); err != nil {
		return err
	}
	for name, source := range uiPrimitives() {
		if err := os.WriteFile(filepath.Join(uiDir, name), []byte(source), 0o644); err != nil {
			return err
		}
	}
	return nil
}
