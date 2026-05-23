//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what EachBlock.Diags 를 파일명 보정 후 수집 슬라이스에 추가
package stml

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func appendEachDiags(eb *EachBlock, file string, out *[]diagnostic.Diagnostic) {
	for _, d := range eb.Diags {
		if d.File == "" {
			d.File = file
		}
		*out = append(*out, d)
	}
}
