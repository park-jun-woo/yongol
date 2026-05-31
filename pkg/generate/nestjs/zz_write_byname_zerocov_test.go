//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteByName_ZeroCov — buildPlansByFeature/writeFeatureModules/writeOneFeature/writeServiceArtifacts 직접 호출

package nestjs

import (
	"os"
	"path/filepath"
	"testing"

	pssac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/types"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildPlansByFeature_Empty_ZeroCov(t *testing.T) {
	// No service funcs → empty map (loop body never runs).
	out := buildPlansByFeature(&yongol.Fullstack{})
	if len(out) != 0 {
		t.Errorf("empty fs should yield empty map, got %v", out)
	}
}

func TestWriteFeatureModules_Empty_ZeroCov(t *testing.T) {
	// Empty plansByFeature → no features written, nil error.
	names, err := writeFeatureModules(map[string][]*ir.ServicePlan{}, t.TempDir(), types.NewRegistry())
	if err != nil {
		t.Fatalf("writeFeatureModules empty: %v", err)
	}
	if len(names) != 0 {
		t.Errorf("expected no feature names, got %v", names)
	}
}

// TestWriteFeatureModules_Integ_ZeroCov drives the full write path
// (writeFeatureModules → writeOneFeature → writeServiceArtifacts) using real
// plans built from the zenflow dummy specs.
func TestWriteFeatureModules_Integ_ZeroCov(t *testing.T) {
	root := "/home/parkjunwoo/.clari/repos/fullend/dummys/zenflow/try-03/specs"
	if _, err := os.Stat(root); err != nil {
		t.Skipf("dummy specs not present: %v", err)
	}
	det, err := yongol.DetectSSOTs(root)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	fs := yongol.ParseAll(root, det)
	if len(fs.ServiceFuncs) == 0 {
		funcs, _ := pssac.ParseDir(filepath.Join(root, "service"))
		fs.ServiceFuncs = funcs
	}
	if len(fs.ServiceFuncs) == 0 {
		t.Skip("no service funcs to build plans from")
	}

	plansByFeature := buildPlansByFeature(fs)
	if len(plansByFeature) == 0 {
		t.Skip("buildPlansByFeature produced no plans")
	}

	reg := types.NewRegistry()

	// Directly exercise writeOneFeature + writeServiceArtifacts for one feature
	// so tsma attributes those names too.
	for feature, plans := range plansByFeature {
		if err := writeOneFeature(feature, plans, t.TempDir(), reg); err != nil {
			t.Fatalf("writeOneFeature(%s): %v", feature, err)
		}
		if len(plans) > 0 {
			if err := writeServiceArtifacts(plans[0], t.TempDir(), reg); err != nil {
				t.Fatalf("writeServiceArtifacts(%s): %v", plans[0].OperationID, err)
			}
		}
		break
	}

	srcDir := t.TempDir()
	names, err := writeFeatureModules(plansByFeature, srcDir, reg)
	if err != nil {
		t.Fatalf("writeFeatureModules: %v", err)
	}
	if len(names) == 0 {
		t.Error("expected at least one feature written")
	}
	// At least one .module.ts must exist.
	var tsFiles int
	_ = filepath.Walk(srcDir, func(p string, info os.FileInfo, e error) error {
		if e == nil && info != nil && !info.IsDir() && filepath.Ext(p) == ".ts" {
			tsFiles++
		}
		return nil
	})
	if tsFiles == 0 {
		t.Error("expected generated .ts files")
	}
}
