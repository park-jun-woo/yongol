//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what filterPagesByDomain — 프론트엔드 디렉토리 기준으로 STML 페이지 필터링
package domain_security

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// filterPagesByDomain returns STML pages whose FileName starts with or is
// relative to the given frontend directory path.
func filterPagesByDomain(pages []stml.PageSpec, frontendDir string) []stml.PageSpec {
	if frontendDir == "" {
		return nil
	}
	// Normalize: ensure no trailing slash for prefix matching.
	prefix := strings.TrimSuffix(frontendDir, "/")

	var result []stml.PageSpec
	for _, p := range pages {
		// FileName may be relative or absolute; match by prefix or base dir.
		dir := filepath.Dir(p.FileName)
		if strings.HasPrefix(p.FileName, prefix) || strings.HasPrefix(dir, prefix) || dir == prefix {
			result = append(result, p)
		}
	}
	return result
}
