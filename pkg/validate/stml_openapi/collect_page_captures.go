//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectPageCaptures — 모든 STML 페이지의 data-capture 바인딩을 (페이지, 바인딩) 쌍으로 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectPageCaptures gathers every parsed data-capture binding across all
// STML pages' action blocks. Syntactically invalid attributes contribute
// nothing (their Captures slice is empty; TM-20 reports the syntax error).
func collectPageCaptures(pages []stml.PageSpec) []pageCapture {
	var out []pageCapture
	for _, p := range pages {
		for _, a := range p.Actions {
			for _, b := range a.Captures {
				out = append(out, pageCapture{File: p.FileName, Bind: b})
			}
		}
	}
	return out
}
