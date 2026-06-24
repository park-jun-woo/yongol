//ff:func feature=gen-react type=generator control=sequence
//ff:what RunTscCheck — 생성된 frontend 를 tsc --noEmit 로 컴파일 게이트 (미해석 시 graceful skip, BUG-137 Phase041)
package react

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunTscCheck spawns `tsc --noEmit` over the generated frontend tree rooted at
// frontendDir so the "generate success = buildable" invariant is enforced for
// the frontend, symmetric to the backend `go build` step. A compile error fails
// generate (the gate); exit 0 passes.
//
// frontendDir is the concrete frontend root: <artifacts>/frontend for a single
// site, or <artifacts>/frontend/<domain> per domain (the caller loops domains).
//
// Graceful skip (warn, not fail) when the toolchain is unresolved: no project
// node_modules (imported packages like react/@tanstack would not resolve — a
// false gate) or no tsc/npx. Auto `npm install` is out of scope.
func RunTscCheck(frontendDir string) error {
	// Prefer the project's own tsc (matches its tsconfig/version); otherwise
	// fall back to env/npx resolution, but only when node_modules exists so
	// package imports resolve.
	var argv []string
	localBin := filepath.Join(frontendDir, "node_modules", ".bin", "tsc")
	if _, err := os.Stat(localBin); err == nil {
		argv = []string{localBin}
	} else {
		if _, err := os.Stat(filepath.Join(frontendDir, "node_modules")); err != nil {
			fmt.Fprintf(os.Stderr, "yongol: skipping frontend tsc gate — node_modules not installed under %s (run `npm install`)\n", frontendDir)
			return nil
		}
		argv = resolveTscArgv()
	}
	if argv == nil {
		fmt.Fprintln(os.Stderr, "yongol: skipping frontend tsc gate — tsc/npx not found")
		return nil
	}

	args := append(append([]string{}, argv[1:]...), "--noEmit")
	cmd := exec.Command(argv[0], args...)
	cmd.Dir = frontendDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("frontend tsc --noEmit failed — generated React does not type-check:\n%s", strings.TrimSpace(string(out)))
	}
	return nil
}
