//ff:func feature=tsx-parser type=command control=sequence
//ff:what swc 외부 바이너리를 spawn 하여 .tsx AST JSON 을 획득
package tsx

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

// swcParserScript is the Node wrapper that calls @swc/core#parseFile and
// writes the AST JSON to stdout. swc's standalone CLI does not expose a
// `parse` subcommand, so the wrapper is the minimal supported entry point.
// It is materialized on disk on first use under $TMPDIR and reused.
const swcParserScript = `
const path = require('path');
const file = process.argv[2];
if (!file) { console.error('usage: swc-parse <file>'); process.exit(2); }
let core;
try {
  core = require('@swc/core');
} catch (e) {
  console.error('yongol: @swc/core not installed. Run one of:');
  console.error('  npm install --save-dev @swc/core');
  console.error('  (at the project root or in YONGOL_SWC_PROJECT_DIR)');
  console.error('underlying error: ' + e.message);
  process.exit(3);
}
core.parseFile(file, {
  syntax: 'typescript',
  tsx: true,
  decorators: false,
  dynamicImport: true,
  comments: false,
}).then(ast => {
  process.stdout.write(JSON.stringify(ast));
}).catch(err => {
  console.error(err && err.message ? err.message : String(err));
  process.exit(1);
});
`

var (
	swcScriptOnce sync.Once
	swcScriptPath string
	swcScriptErr  error
)

// ensureSwcScript writes the bundled parser wrapper into a deterministic
// temp file and returns its path. The script is tiny and stateless, so
// reuse across invocations is safe.
func ensureSwcScript() (string, error) {
	swcScriptOnce.Do(func() {
		dir := filepath.Join(os.TempDir(), "yongol-tsx-parser")
		if err := os.MkdirAll(dir, 0o755); err != nil {
			swcScriptErr = fmt.Errorf("create swc script dir: %w", err)
			return
		}
		path := filepath.Join(dir, "parse.cjs")
		if err := os.WriteFile(path, []byte(swcParserScript), 0o644); err != nil {
			swcScriptErr = fmt.Errorf("write swc script: %w", err)
			return
		}
		swcScriptPath = path
	})
	return swcScriptPath, swcScriptErr
}

// swcProjectDir returns the directory node should run from so that a local
// @swc/core install is discovered. The resolution order is:
//
//  1. $YONGOL_SWC_PROJECT_DIR (explicit override)
//  2. ascending from the target file: nearest ancestor with node_modules/@swc/core
//  3. ascending from the target file: nearest ancestor with package.json
//  4. the file's own directory
//
// The caller is responsible for a clear diagnostic when none of the above
// resolves to a working install; swcRunnerNotice is the canonical message.
func swcProjectDir(target string) string {
	if v := os.Getenv("YONGOL_SWC_PROJECT_DIR"); v != "" {
		return v
	}
	abs, err := filepath.Abs(target)
	if err != nil {
		abs = target
	}
	dir := filepath.Dir(abs)
	// Walk up to filesystem root.
	for {
		if st, err := os.Stat(filepath.Join(dir, "node_modules", "@swc", "core")); err == nil && st.IsDir() {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	// Second pass: nearest package.json.
	dir = filepath.Dir(abs)
	for {
		if _, err := os.Stat(filepath.Join(dir, "package.json")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return filepath.Dir(abs)
}

// swcNodeBin resolves the node binary. $YONGOL_NODE wins; otherwise PATH.
func swcNodeBin() (string, error) {
	if v := os.Getenv("YONGOL_NODE"); v != "" {
		return v, nil
	}
	if p, err := exec.LookPath("node"); err == nil {
		return p, nil
	}
	return "", fmt.Errorf("node not found on PATH (set YONGOL_NODE to override)")
}

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

// swcRunnerNotice is the fail-fast installation hint surfaced when the swc
// toolchain is absent. Kept as a single const so CLI integrations can wrap
// it uniformly.
const swcRunnerNotice = "install Node.js (>=18) and run `npm install --save-dev @swc/core` in your frontend project (or set YONGOL_SWC_PROJECT_DIR)"
