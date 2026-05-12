//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what 모든 LayoutSpec에 대해 레이아웃 TSX 파일을 생성한다

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// writeLayoutsTSX emits layout TSX files for all provided LayoutSpecs.
func writeLayoutsTSX(srcDir string, layouts []stml.LayoutSpec) error {
	for _, l := range layouts {
		if err := writeLayoutTSX(srcDir, l); err != nil {
			return err
		}
	}
	return nil
}
