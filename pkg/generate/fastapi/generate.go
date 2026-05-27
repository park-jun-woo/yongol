//ff:func feature=gen-fastapi type=command control=sequence
//ff:what Generate — FastAPI 백엔드 코드 생성 (IR → Python 렌더링 파이프라인)

package fastapi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/fastapi/types"
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces FastAPI backend artifacts from a parsed Fullstack.
// Pipeline:
//  1. Build IR plans (ServicePlan, BootPlan, MiddlewarePlan, InfraPlan)
//  2. Render scaffold (pyproject.toml, requirements.txt, .env.example)
//  3. Render SQLAlchemy models from DDL tables
//  4. Generate infrastructure stubs (event_bus, authz, external packages)
//  5. For each service func, render router + service
//  6. Render main.py + config.py + database.py
//  7. Write all files to dir
func Generate(fs *yongol.Fullstack, dir string) error {
	psVal := prepared.New(fs)
	ps := &psVal
	reg := types.NewRegistry()

	backendDir := filepath.Join(dir, "backend")
	appDir := filepath.Join(backendDir, "app")
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return fmt.Errorf("mkdir app: %w", err)
	}

	bootPlan := ir.BuildBootPlan(fs, ps, "fastapi")
	_ = ir.BuildMiddlewarePlan(fs, ps)
	_ = ir.BuildInfraPlan(fs, ps)

	plansByFeature := buildPlansByFeature(fs)
	projectID := resolveProjectID(fs)

	if err := writeScaffold(backendDir, projectID); err != nil {
		return fmt.Errorf("scaffold: %w", err)
	}
	if err := writeModels(fs, appDir); err != nil {
		return fmt.Errorf("models: %w", err)
	}
	if err := writeDependencies(appDir); err != nil {
		return fmt.Errorf("dependencies: %w", err)
	}

	// Generate event bus infrastructure when @publish exists.
	if hasPublishPlans(plansByFeature) {
		if err := writeEventBus(appDir); err != nil {
			return fmt.Errorf("event bus: %w", err)
		}
	}

	// Generate authz infrastructure when @auth exists.
	if hasAuthPlans(plansByFeature) {
		if err := writeAuthz(appDir); err != nil {
			return fmt.Errorf("authz: %w", err)
		}
	}

	// Generate stub modules for external packages referenced by @call/@eval.
	extPkgs := collectExternalPackages(plansByFeature)
	if len(extPkgs) > 0 {
		if err := writeFuncStubs(appDir, extPkgs); err != nil {
			return fmt.Errorf("func stubs: %w", err)
		}
	}

	featureNames, err := writeFeatureModules(plansByFeature, appDir, reg)
	if err != nil {
		return err
	}
	if err := writeInitFiles(appDir, featureNames); err != nil {
		return fmt.Errorf("init files: %w", err)
	}
	return writeBootFiles(appDir, bootPlan, featureNames)
}
