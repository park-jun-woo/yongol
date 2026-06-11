//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what LinkRef에서 react-router Link의 to 속성 문자열을 생성한다 (매핑 치환, optional 미매핑 생략)
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// LinkToAttr builds the `to` attribute of the emitted <Link> by
// substituting the link's param sources into the target page's resolved
// route pattern. Optional segments (":Name?") are filled only when mapped
// and omitted otherwise; unmapped required segments cannot survive
// validation (TM-32) and are omitted defensively. Without any
// interpolation the attribute is a plain string, otherwise a template
// literal. An empty TargetPattern falls back to "/<page-name>".
func LinkToAttr(lr stmlparser.LinkRef) string {
	pattern := lr.TargetPattern
	if pattern == "" {
		pattern = "/" + lr.TargetPage
	}
	sources := linkSegmentSources(lr)
	var parts []string
	hasExpr := false
	for _, seg := range strings.Split(strings.TrimPrefix(pattern, "/"), "/") {
		if seg == "" {
			continue
		}
		if !strings.HasPrefix(seg, ":") {
			parts = append(parts, seg)
			continue
		}
		name := strings.TrimSuffix(strings.TrimPrefix(seg, ":"), "?")
		src, ok := sources[name]
		if !ok {
			continue
		}
		parts = append(parts, "${"+linkParamExpr(src)+"}")
		hasExpr = true
	}
	path := "/" + strings.Join(parts, "/")
	if hasExpr {
		return fmt.Sprintf("to={`%s`}", path)
	}
	return fmt.Sprintf("to=%q", path)
}
