//ff:func feature=validate type=test-helper control=sequence topic=manifest-structural
//ff:what fsWithDomains — domains 맵만 채운 Fullstack 픽스처 생성 헬퍼

package manifest

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// fsWithDomains builds a Fullstack whose manifest carries only the given
// domains map — the minimal fixture the C-12~C-17 rules read.
func fsWithDomains(d map[string]pmanifest.DomainConfig) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Domains: d}}
}
