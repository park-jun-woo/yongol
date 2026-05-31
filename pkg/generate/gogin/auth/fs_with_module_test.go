//ff:func feature=gen-gogin type=test-helper control=sequence
//ff:what authFSWithModule — backend.Module 이 설정된 Fullstack 생성 헬퍼
package auth

import (
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// authFSWithModule returns a Fullstack whose backend module path is set.
func authFSWithModule() *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Module: "github.com/test/app"},
		},
	}
}
