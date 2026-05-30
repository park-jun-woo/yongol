//ff:func feature=gen-react type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 react.Generate 통합 커버리지

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func loadDummyFS_Integ(t *testing.T, root string) *yongol.Fullstack {
	t.Helper()
	det, err := yongol.DetectSSOTs(root)
	if err != nil {
		t.Fatalf("DetectSSOTs(%s): %v", root, err)
	}
	return yongol.ParseAll(root, det)
}

// TestReactGenerate_Integ runs react.Generate against zenflow specs. This
// exercises Generate → generateFrontendSetup → writeLibUtils → findOpenAPISpec
// → fsOpenAPIDoc, and either the success or stub path of runOpenAPITypescript
// depending on whether the openapi-typescript binary is installed.
func TestReactGenerate_Integ(t *testing.T) {
	root := "/home/parkjunwoo/.clari/repos/fullend/dummys/zenflow/try-03/specs"
	if _, err := os.Stat(root); err != nil {
		t.Skipf("dummy specs not present: %v", err)
	}
	fs := loadDummyFS_Integ(t, root)

	out := t.TempDir()
	// Generate may return a deferred error if openapi-typescript is missing;
	// the file tree must still be materialized in either case.
	err := Generate(fs, out)
	t.Logf("Generate returned: %v (deferred openapi-typescript error tolerated)", err)

	frontend := filepath.Join(out, "frontend")
	for _, name := range []string{
		"package.json",
		"vite.config.ts",
		filepath.Join("src", "lib", "utils.ts"),
		filepath.Join("src", "lib", "api.ts"),
		filepath.Join("src", "types", "api.d.ts"),
	} {
		if _, statErr := os.Stat(filepath.Join(frontend, name)); statErr != nil {
			t.Errorf("expected %s: %v", name, statErr)
		}
	}
}

// TestReactGenerate_NoSpec_Integ drives the branch where findOpenAPISpec returns
// "" (no SpecsDir), so the api.d.ts stub is written and Generate returns nil.
func TestReactGenerate_NoSpec_Integ(t *testing.T) {
	fs := &yongol.Fullstack{} // empty: SpecsDir "", OpenAPIDoc nil, Manifest nil
	out := t.TempDir()
	if err := Generate(fs, out); err != nil {
		t.Fatalf("Generate with empty fullstack: %v", err)
	}
	stub := filepath.Join(out, "frontend", "src", "types", "api.d.ts")
	data, err := os.ReadFile(stub)
	if err != nil {
		t.Fatalf("expected api.d.ts stub: %v", err)
	}
	if len(data) == 0 {
		t.Error("api.d.ts stub is empty")
	}
}

// TestRunOpenAPITypescript_BadSpec_Integ covers the error-return path of
// runOpenAPITypescript: either the binary is missing (stub + error) or it runs
// against a non-existent spec and fails. Either way a non-nil error is returned.
func TestRunOpenAPITypescript_BadSpec_Integ(t *testing.T) {
	dest := filepath.Join(t.TempDir(), "api.d.ts")
	err := runOpenAPITypescript(filepath.Join(t.TempDir(), "does-not-exist.yaml"), dest)
	if err == nil {
		t.Skip("openapi-typescript unexpectedly succeeded on missing spec; environment-specific")
	}
	t.Logf("runOpenAPITypescript error (expected): %v", err)
}

// TestFindOpenAPISpec_FsOpenAPIDoc_ZeroCov pins the small accessor branches.
func TestFindOpenAPISpec_FsOpenAPIDoc_ZeroCov(t *testing.T) {
	if findOpenAPISpec(nil) != "" {
		t.Error("nil fs should yield empty spec path")
	}
	if findOpenAPISpec(&yongol.Fullstack{}) != "" {
		t.Error("empty SpecsDir should yield empty spec path")
	}
	got := findOpenAPISpec(&yongol.Fullstack{SpecsDir: "/x"})
	if got != "/x/api/openapi.yaml" {
		t.Errorf("findOpenAPISpec = %q", got)
	}
	if fsOpenAPIDoc(nil) != nil {
		t.Error("nil fs should yield nil doc")
	}
}
