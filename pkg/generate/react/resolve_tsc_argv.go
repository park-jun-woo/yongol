//ff:func feature=gen-react type=loader control=iteration dimension=1
//ff:what resolveTscArgv — tsc 호출용 argv 해결 (project node_modules → npx -p typescript), 없으면 nil
package react

import (
	"os"
	"os/exec"
	"path/filepath"
)

// resolveTscArgv returns argv for invoking tsc, or nil when none is found.
// Order mirrors resolveOpenapiTsBinary: project node_modules first, then a
// pinned `npx -p typescript tsc` (PATH `tsc` is skipped — environments may
// shadow it with an unrelated binary). A nil return lets callers gracefully
// skip the frontend tsc gate instead of failing generate (BUG-137 Phase041).
func resolveTscArgv() []string {
	for _, dir := range []string{
		os.Getenv("YONGOL_OPENAPI_TS_PROJECT_DIR"),
		os.Getenv("YONGOL_SWC_PROJECT_DIR"),
	} {
		if dir == "" {
			continue
		}
		local := filepath.Join(dir, "node_modules", ".bin", "tsc")
		if _, err := os.Stat(local); err == nil {
			return []string{local}
		}
	}
	if _, err := exec.LookPath("npx"); err == nil {
		return []string{"npx", "--yes", "-p", "typescript", "tsc"}
	}
	return nil
}
