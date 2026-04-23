//ff:func feature=tsx-parser type=loader control=sequence
//ff:what swcNodeBin — node 바이너리 경로 결정 ($YONGOL_NODE → PATH)

package tsx

import (
	"fmt"
	"os"
	"os/exec"
)

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
