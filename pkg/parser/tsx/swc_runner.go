//ff:func feature=tsx-parser type=command control=sequence
//ff:what runSwcParse — swc 외부 바이너리를 spawn 하여 .tsx AST JSON 을 획득

package tsx

import (
	"bytes"
	"fmt"
	"os/exec"
	"os"
	"path/filepath"
	"runtime"
)

// runSwcParse spawns node with the bundled wrapper and returns the raw AST
// JSON bytes. Returns a descriptive error if node is missing, @swc/core is
// not installed at the resolved project dir, or the wrapper fails.
func runSwcParse(file string) ([]byte, error) {
	script, err := ensureSwcScript()
	if err != nil {
		return nil, err
	}
	node, err := swcNodeBin()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, swcRunnerNotice)
	}
	abs, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}
	projDir := swcProjectDir(abs)
	cmd := exec.Command(node, script, abs)
	cmd.Dir = projDir
	// Node resolves `require('@swc/core')` relative to the script file, not
	// cwd. Exposing the project's node_modules via NODE_PATH makes the
	// resolver fall back to it. This is the documented Node mechanism for
	// out-of-tree scripts.
	cmd.Env = append(os.Environ(), "NODE_PATH="+filepath.Join(projDir, "node_modules"))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf(
			"swc parse failed (%s, proj=%s, runtime=%s): %v: %s",
			file, projDir, runtime.GOOS, err, stderr.String(),
		)
	}
	return stdout.Bytes(), nil
}
