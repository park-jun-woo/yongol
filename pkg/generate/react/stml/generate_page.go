//ff:func feature=stml-gen type=generator control=sequence
//ff:what 기본 Target으로 단일 페이지의 소스 코드를 생성한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// GeneratePage generates source code for a single page using the default target.
func GeneratePage(page stmlparser.PageSpec, specsDir string, opts ...GenerateOptions) string {
	opt := DefaultOptions()
	if len(opts) > 0 {
		opt = mergeOpt(opt, opts[0])
	}
	return DefaultTarget().GeneratePage(page, specsDir, opt)
}
