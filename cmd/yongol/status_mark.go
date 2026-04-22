//ff:func feature=cli type=util control=selection
//ff:what statusMark — validate.Status를 ✓/✗/-/? 기호로 변환
package main

import "github.com/park-jun-woo/yongol/pkg/validate"

func statusMark(st validate.Status) string {
	switch st {
	case validate.StatusPass:
		return "✓"
	case validate.StatusFail:
		return "✗"
	case validate.StatusSkip:
		return "-"
	case validate.StatusMissing:
		return "?"
	}
	return " "
}
