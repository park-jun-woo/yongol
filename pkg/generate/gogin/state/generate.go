//ff:func feature=gen-gogin type=command control=iteration dimension=1
//ff:what Generate — StateDiagram 순회하여 statemachine Go 파일 생성 (1파일 1func 분리)

package state

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces three files per StateDiagram under
// internal/statemachine/: <id>.go (Transitions var), <id>_can_transition.go,
// <id>_next_state.go. Skipped when no StateDiagrams exist.
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	if len(fs.StateDiagrams) == 0 {
		return nil
	}
	dir := filepath.Join(artifactsDir, "backend", "internal", "statemachine")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("mkdir statemachine: %w", err)
	}
	for _, d := range fs.StateDiagrams {
		transMap := buildTransitionMap(d)
		if err := writeStateFile(dir, d.ID, transMap); err != nil {
			return fmt.Errorf("statemachine %s: %w", d.ID, err)
		}
	}
	return nil
}
