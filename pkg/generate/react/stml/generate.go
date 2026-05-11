//ff:func feature=stml-gen type=generator control=sequence topic=output
//ff:what 기본 Target으로 페이지 목록의 프레임워크별 파일을 생성한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// Generate produces framework-specific files using the default target.
func Generate(pages []stmlparser.PageSpec, specsDir, outDir string, opts ...GenerateOptions) (*GenerateResult, error) {
	return GenerateWith(DefaultTarget(), pages, specsDir, outDir, opts...)
}
