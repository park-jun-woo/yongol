//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeOneFeature — 단일 feature 의 service+controller+module 파일 기록

package nestjs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/ssac"
)

// writeOneFeature writes all artifacts for a single feature.
func writeOneFeature(feature string, plans []*ir.ServicePlan, srcDir string, reg ir.TypeRegistry) error {
	featureDir := filepath.Join(srcDir, feature)
	if err := os.MkdirAll(featureDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", feature, err)
	}
	for _, plan := range plans {
		if err := writeServiceArtifacts(plan, featureDir, reg); err != nil {
			return err
		}
	}
	modContent, err := ssac.RenderModule(feature, plans)
	if err != nil {
		return fmt.Errorf("render module %s: %w", feature, err)
	}
	return os.WriteFile(filepath.Join(featureDir, feature+".module.ts"), []byte(modContent), 0o644)
}
