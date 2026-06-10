//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what parseParamBinds — "src -> Segment, ..." 매핑 문법의 공유 코어 (link/redirect params가 소스 검사만 달리해 재사용)
package stml

import (
	"fmt"
	"strings"
)

// parseParamBinds is the shared core of the value-based "a -> b" mapping
// grammar (ParseLinkParams / ParseRedirectParams). label names the
// attribute family in error messages ("link param" / "redirect param");
// checkSource validates the left-hand side — the only part where the two
// attributes differ (link sources are item.*/route.*, redirect sources are
// bare 2xx respFields or route.*). The "-> <SegmentName>" part may be
// elided; the per-rule validators (TM-32 / TM-33) then require the target
// route to have exactly one required segment.
func parseParamBinds(raw, label string, checkSource func(source string) error) ([]LinkParamBind, error) {
	var out []LinkParamBind
	for _, seg := range strings.Split(raw, ",") {
		seg = strings.TrimSpace(seg)
		if seg == "" {
			return nil, fmt.Errorf("empty %s binding; expected \"<source> -> <SegmentName>\"", label)
		}
		parts := strings.Split(seg, "->")
		if len(parts) > 2 {
			return nil, fmt.Errorf("%s binding %q must be \"<source> -> <SegmentName>\"", label, seg)
		}
		source := strings.TrimSpace(parts[0])
		if err := checkSource(source); err != nil {
			return nil, err
		}
		segment := ""
		if len(parts) == 2 {
			segment = strings.TrimSpace(parts[1])
		}
		if len(parts) == 2 && segment == "" {
			return nil, fmt.Errorf("%s binding %q has an empty segment name", label, seg)
		}
		out = append(out, LinkParamBind{Source: source, Segment: segment})
	}
	return out, nil
}
