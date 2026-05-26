//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeServiceArtifacts — 단일 ServicePlan 의 .service.ts + .controller.ts 파일 기록

package nestjs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/ssac"
)

// writeServiceArtifacts writes the .service.ts and .controller.ts for one plan.
func writeServiceArtifacts(plan *ir.ServicePlan, featureDir string, reg ir.TypeRegistry) error {
	baseName := nestLcFirst(plan.OperationID)
	svcContent, err := ssac.RenderService(plan, reg)
	if err != nil {
		return fmt.Errorf("render service %s: %w", plan.OperationID, err)
	}
	if err := os.WriteFile(filepath.Join(featureDir, baseName+".service.ts"), []byte(svcContent), 0o644); err != nil {
		return fmt.Errorf("write service %s: %w", plan.OperationID, err)
	}
	ctrlContent, err := ssac.RenderController(plan)
	if err != nil {
		return fmt.Errorf("render controller %s: %w", plan.OperationID, err)
	}
	return os.WriteFile(filepath.Join(featureDir, baseName+".controller.ts"), []byte(ctrlContent), 0o644)
}
