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
//  4. For each service func, render router + service
//  5. Render main.py + config.py + database.py
//  6. Write all files to dir
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
	featureNames, err := writeFeatureModules(plansByFeature, appDir, reg)
	if err != nil {
		return err
	}
	if err := writeInitFiles(appDir, featureNames); err != nil {
		return fmt.Errorf("init files: %w", err)
	}
	return writeBootFiles(appDir, bootPlan, featureNames)
}
