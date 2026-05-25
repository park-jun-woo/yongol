//ff:func feature=gen-gogin type=generator control=iteration dimension=1 topic=auth-refresh
//ff:what emitAuthWrapper — ssac/pkg/auth.RefreshStore adapter 를 6파일로 분할 emit (1 file 1 func/type)

package infra

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/ffhash"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// emitAuthWrapper writes adapter files under arts/backend/internal/infra/auth/
// with one file per method (filefunc F3) plus the type + constructor file.
func emitAuthWrapper(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error {
	insertPort := portByName(active, "RefreshTokenInsert")
	consumePort := portByName(active, "RefreshTokenConsume")
	checkReusePort := portByName(active, "RefreshTokenCheckReuse")
	revokePort := portByName(active, "RefreshTokenRevoke")
	revokeAllPort := portByName(active, "RefreshTokenRevokeAll")
	if insertPort == nil || consumePort == nil || checkReusePort == nil || revokePort == nil || revokeAllPort == nil {
		return fmt.Errorf("auth: interface.yaml missing one of RefreshTokenInsert/Consume/CheckReuse/Revoke/RevokeAll (active ports: %d)", len(active))
	}

	dir := filepath.Join(artifactsDir, "backend", "internal", "infra", iface.Package)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	files := map[string]string{
		"postgres.go":            fmt.Sprintf(authWrapperTypeTemplate, modulePath),
		"postgres_new.go":        fmt.Sprintf(authWrapperNewPostgresTemplate, modulePath),
		"postgres_create.go":     fmt.Sprintf(authWrapperCreateTemplate, modulePath, insertPort.Name),
		"postgres_consume.go":    fmt.Sprintf(authWrapperConsumeTemplate, modulePath, consumePort.Name, checkReusePort.Name),
		"postgres_revoke.go":     fmt.Sprintf(authWrapperRevokeTemplate, modulePath, revokePort.Name),
		"postgres_revoke_all.go": fmt.Sprintf(authWrapperRevokeAllTemplate, modulePath, revokeAllPort.Name),
	}
	for name, content := range files {
		data := ffhash.InjectCheckedLine([]byte(content))
		if err := os.WriteFile(filepath.Join(dir, name), data, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", name, err)
		}
	}
	return nil
}
