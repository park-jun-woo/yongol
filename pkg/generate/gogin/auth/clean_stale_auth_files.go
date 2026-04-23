//ff:func feature=gen-gogin type=util control=sequence
//ff:what cleanStaleAuthFiles — internal/auth/ 디렉토리 전체를 제거한다 (Phase001 UserClaimUnification)

package auth

import (
	"os"
	"path/filepath"
)

// cleanStaleAuthFiles removes the entire `backend/internal/auth/` directory
// from artifactsDir. Phase001 UserClaimUnification collapsed everything
// auth-related into ssac/pkg/auth + model.UserClaim, so nothing under
// internal/auth/ survives regeneration; wiping the directory wholesale
// avoids leftover files (claim.go / reexport.go / issue_token.go /
// refresh_token.go / verify_token.go / refresh_store.go / refresh_handler.go)
// from earlier yongol versions polluting the generated module.
//
// Missing directories are ignored — a brand-new artifactsDir never had an
// internal/auth/ tree and should not fail clean-up.
func cleanStaleAuthFiles(artifactsDir string) error {
	authDir := filepath.Join(artifactsDir, "backend", "internal", "auth")
	if err := os.RemoveAll(authDir); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
