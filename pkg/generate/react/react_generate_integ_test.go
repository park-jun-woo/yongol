//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 react.Generate 통합 커버리지
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReactGenerate_Integ(t *testing.T) {
	root := "/home/parkjunwoo/.clari/repos/fullend/examples/zenflow/try-03/specs"
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
