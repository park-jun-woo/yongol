//ff:func feature=gen-nestjs type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 nestjs.Generate 통합 커버리지

package nestjs

import (
	"os"
	"path/filepath"
	"testing"

	pssac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// loadDummyFS_Integ builds a real *yongol.Fullstack from a dummy specs dir.
// ServiceFuncs is force-populated from ssac.ParseDir because some dummy specs
// carry known parse diagnostics that would otherwise gate ServiceFuncs to nil
// inside ParseAll — for generator integration coverage we want the funcs.
func loadDummyFS_Integ(t *testing.T, root string) *yongol.Fullstack {
	t.Helper()
	det, err := yongol.DetectSSOTs(root)
	if err != nil {
		t.Fatalf("DetectSSOTs(%s): %v", root, err)
	}
	fs := yongol.ParseAll(root, det)
	if len(fs.ServiceFuncs) == 0 {
		funcs, _ := pssac.ParseDir(filepath.Join(root, "service"))
		fs.ServiceFuncs = funcs
	}
	return fs
}

// TestNestjsGenerate_Integ runs nestjs.Generate against the zenflow dummy specs,
// which exercises buildPlansByFeature, writeFeatureModules, writeOneFeature and
// writeServiceArtifacts in one orchestrated pass.
func TestNestjsGenerate_Integ(t *testing.T) {
	root := "/home/parkjunwoo/.clari/repos/fullend/dummys/zenflow/try-03/specs"
	if _, err := os.Stat(root); err != nil {
		t.Skipf("dummy specs not present: %v", err)
	}
	fs := loadDummyFS_Integ(t, root)
	if len(fs.ServiceFuncs) == 0 {
		t.Fatal("expected ServiceFuncs from dummy specs")
	}

	out := t.TempDir()
	if err := Generate(fs, out); err != nil {
		t.Fatalf("Generate: %v", err)
	}

	// Feature module + service artifacts should exist somewhere under the tree.
	var tsFiles int
	_ = filepath.Walk(out, func(p string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() && filepath.Ext(p) == ".ts" {
			tsFiles++
		}
		return nil
	})
	if tsFiles == 0 {
		t.Error("expected generated .ts files under output dir")
	}
}
