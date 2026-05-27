//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what writeOneFeature — 단일 feature 의 service+router 파일 기록

package fastapi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/fastapi/ssac"
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeOneFeature writes all artifacts for a single feature.
func writeOneFeature(feature string, plans []*ir.ServicePlan, appDir string, reg ir.TypeRegistry) error {
	// Write services
	servicesDir := filepath.Join(appDir, "services")
	if err := os.MkdirAll(servicesDir, 0o755); err != nil {
		return fmt.Errorf("mkdir services: %w", err)
	}

	// Write consolidated imports once, then each function body.
	imports := ssac.WriteFeatureImports(plans, feature)
	var svcContent string
	svcContent += imports
	for _, plan := range plans {
		content, err := ssac.RenderService(plan, reg)
		if err != nil {
			return fmt.Errorf("render service %s: %w", plan.OperationID, err)
		}
		svcContent += content + "\n\n"
	}
	if err := os.WriteFile(filepath.Join(servicesDir, feature+".py"), []byte(svcContent), 0o644); err != nil {
		return fmt.Errorf("write service %s: %w", feature, err)
	}

	// Write router
	routersDir := filepath.Join(appDir, "routers")
	if err := os.MkdirAll(routersDir, 0o755); err != nil {
		return fmt.Errorf("mkdir routers: %w", err)
	}

	routerContent, err := ssac.RenderRouter(feature, plans)
	if err != nil {
		return fmt.Errorf("render router %s: %w", feature, err)
	}
	return os.WriteFile(filepath.Join(routersDir, feature+".py"), []byte(routerContent), 0o644)
}
