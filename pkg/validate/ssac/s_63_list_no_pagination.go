//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-63 — @get []T list 엔드포인트인데 pagination params 없고 @no-pagination 없음 (WARNING)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// paginationKeys는 pagination을 나타내는 Inputs 키 목록이다.
var paginationKeys = map[string]bool{
	"Page":    true,
	"PerPage": true,
	"Cursor":  true,
}

// s63ListNoPagination validates S-63: a @get sequence whose Result.Type starts
// with "[]" (list) should have at least one pagination key (Page, PerPage,
// Cursor) in its Inputs. Emits WARNING if missing and the function does not
// carry the // @no-pagination annotation.
func s63ListNoPagination(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.NoPagination {
			continue
		}
		for _, seq := range fn.Sequences {
			if seq.Type != "get" {
				continue
			}
			if seq.Result == nil {
				continue
			}
			if !strings.HasPrefix(seq.Result.Type, "[]") {
				continue
			}
			if hasPaginationKey(seq.Inputs) {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: "[S-63] @get []T list endpoint has no pagination (missing @no-pagination)",
				Advice:  "pagination params (page/per_page/cursor) 를 추가하거나, 의도적으로 전체 목록을 반환한다면 // @no-pagination 을 붙이세요",
			})
			break // 함수당 한 번만 경고
		}
	}
	return diags
}
