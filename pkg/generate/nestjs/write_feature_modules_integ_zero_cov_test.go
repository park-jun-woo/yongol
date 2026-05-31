//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteByName_ZeroCov — buildPlansByFeature/writeFeatureModules/writeOneFeature/writeServiceArtifacts 직접 호출
package nestjs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/types"
	pssac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
	exerciseOneFeature(t, plansByFeature, reg)

	srcDir := t.TempDir()
	names, err := writeFeatureModules(plansByFeature, srcDir, reg)
	if err != nil {
		t.Fatalf("writeFeatureModules: %v", err)
	}
	if len(names) == 0 {
		t.Error("expected at least one feature written")
	}
	// At least one .module.ts must exist.
	if countTSFiles(srcDir) == 0 {
		t.Error("expected generated .ts files")
	}
}
