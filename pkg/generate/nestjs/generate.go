//ff:func feature=gen-nestjs type=command control=sequence
//ff:what Generate — NestJS 백엔드 코드 생성 (IR → TypeScript 렌더링 파이프라인, 인프라 모듈 등록)

package nestjs

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/types"
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces NestJS backend artifacts from a parsed Fullstack.
// Pipeline:
//  1. Build IR plans (ServicePlan, BootPlan, MiddlewarePlan, InfraPlan)
//  2. Render scaffold (package.json, tsconfig.json, nest-cli.json)
//  3. Render Prisma schema from DDL tables
//  4. For each service func, render controller + service + module
//  5. Generate infrastructure stubs (queue, authz, external packages)
//  6. Render main.ts + app.module.ts (with infra + feature modules)
//  7. Write all files to dir
func Generate(fs *yongol.Fullstack, dir string) error {
	psVal := prepared.New(fs)
	ps := &psVal
	reg := types.NewRegistry()

	backendDir := filepath.Join(dir, "backend")
	srcDir := filepath.Join(backendDir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		return fmt.Errorf("mkdir src: %w", err)
	}

	bootPlan := ir.BuildBootPlan(fs, ps, "nestjs")
	_ = ir.BuildMiddlewarePlan(fs, ps)
	_ = ir.BuildInfraPlan(fs, ps)

	plansByFeature := buildPlansByFeature(fs)
	projectID := resolveProjectID(fs)

	if err := writeScaffold(backendDir, projectID); err != nil {
		return fmt.Errorf("scaffold: %w", err)
	}
	if err := writePrismaSchema(fs, backendDir); err != nil {
		return fmt.Errorf("prisma schema: %w", err)
	}
	if err := writePrismaModule(srcDir); err != nil {
		return fmt.Errorf("prisma module: %w", err)
	}

	// Collect infrastructure module names for app.module.ts registration.
	// Only truly global modules go here (queue, authz). External package
	// stubs that also have feature modules are excluded to avoid duplicates.
	var infraModules []string

	// Generate queue infrastructure when @publish exists.
	if hasPublishPlans(plansByFeature) {
		if err := writeQueueModule(srcDir); err != nil {
			return fmt.Errorf("queue module: %w", err)
		}
		infraModules = append(infraModules, "queue")
	}

	// Generate authz infrastructure when @auth exists.
	if hasAuthPlans(plansByFeature) {
		if err := writeAuthzModule(srcDir); err != nil {
			return fmt.Errorf("authz module: %w", err)
		}
		infraModules = append(infraModules, "authz")
	}

	// Generate stub services for external packages referenced by @call/@eval.
	extPkgs := collectExternalPackages(plansByFeature)
	if len(extPkgs) > 0 {
		if err := writeFuncStubs(srcDir, extPkgs); err != nil {
			return fmt.Errorf("func stubs: %w", err)
		}
		// Only register as infra module if not already a feature module.
		for _, ep := range extPkgs {
			if _, isFeature := plansByFeature[ep.Name]; !isFeature {
				infraModules = append(infraModules, ep.Name)
			}
		}
	}

	sort.Strings(infraModules)

	featureNames, err := writeFeatureModules(plansByFeature, srcDir, reg)
	if err != nil {
		return err
	}
	sort.Strings(featureNames)
	return writeBootFiles(srcDir, bootPlan, featureNames, infraModules)
}
