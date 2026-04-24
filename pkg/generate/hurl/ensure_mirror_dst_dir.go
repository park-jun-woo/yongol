//ff:func feature=gen-hurl type=util control=sequence
//ff:what ensureMirrorDstDir — 최초 user hurl 매치 시 dstDir 를 lazy 생성

package hurl

import (
	"fmt"
	"os"
)

// ensureMirrorDstDir lazily creates dstDir when the first user hurl match
// appears. Returns the updated "has the directory been created?" flag
// (idempotent: once true, subsequent calls short-circuit without an
// os.MkdirAll). Extracted from mirrorUserHurlEntries so the iteration
// body stays within the Q1 depth=2 budget.
func ensureMirrorDstDir(dstDir string, created bool) (bool, error) {
	if created {
		return true, nil
	}
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return false, fmt.Errorf("mkdir %s: %w", dstDir, err)
	}
	return true, nil
}
