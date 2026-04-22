//ff:func feature=generate type=util control=selection
//ff:what runBackend — BackendType에 따라 해당 생성기 실행
package generate

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin"
)

func runBackend(fs *yongol.Fullstack, artifactsDir string, backend BackendType) error {
	switch backend {
	case GoGin:
		return gogin.Generate(fs, artifactsDir)
	default:
		return fmt.Errorf("unknown backend %q", backend)
	}
}
