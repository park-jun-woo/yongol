//ff:func feature=gen-gogin type=util control=sequence
//ff:what runGo — `go <args>` 를 실행하고 실패 시 stderr 를 포함해 래핑

package gogin

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// runGo executes `go <args...>` in dir, capturing stderr for error surfacing.
// Any non-zero exit wraps with the (truncated) stderr so callers see
// the concrete failure (network error, unresolved module, malformed flag)
// rather than a bare exit status.
func runGo(dir string, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go %s failed: %w: %s",
			strings.Join(args, " "), err,
			truncateStderr(stderr.String(), goModTidyStderrLimit))
	}
	return nil
}
