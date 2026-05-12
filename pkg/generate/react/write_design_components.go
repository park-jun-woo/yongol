//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what DESIGN.md 컴포넌트 정의에서 TSX 파일을 생성한다

package react

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// writeDesignComponents emits one TSX file per DESIGN.md component definition.
func writeDesignComponents(uiDir string, comps map[string]design.ComponentToken) error {
	for name, tok := range comps {
		src := renderComponentTSX(name, tok)
		fileName := name + ".tsx"
		if err := os.WriteFile(filepath.Join(uiDir, fileName), []byte(src), 0o644); err != nil {
			return err
		}
	}
	return nil
}
