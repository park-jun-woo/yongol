//ff:func feature=stml-parse type=parser control=sequence
//ff:what ParseLinkParams — data-link-params 값("src -> Segment, ...")을 LinkParamBind 목록으로 파싱 (TM-32 구문 근거)
package stml

import (
	"fmt"
	"strings"
)

// ParseLinkParams parses a data-link-params attribute value of the form
// "<source> -> <SegmentName>[, <source> -> <SegmentName>...]" into
// LinkParamBind entries — the same value-based "a -> b" grammar as
// ParseCapture (a kebab-case attribute suffix cannot spell a PascalCase
// segment name). The "-> <SegmentName>" part may be elided; TM-32 then
// requires the target route to have exactly one required segment. Sources
// are restricted to "item.<Field>" and "route.<Name>". Any format
// violation returns an error; TM-32 re-parses the raw attribute at
// validate time to surface it as a diagnostic (the ParseCapture / TM-20
// split). The grammar core is shared with ParseRedirectParams
// (parseParamBinds) — only the source check differs.
func ParseLinkParams(raw string) ([]LinkParamBind, error) {
	return parseParamBinds(raw, "link param", func(source string) error {
		if !strings.HasPrefix(source, "item.") && !strings.HasPrefix(source, "route.") {
			return fmt.Errorf("link param source %q must be \"item.<Field>\" or \"route.<Name>\"", source)
		}
		if source == "item." || source == "route." {
			return fmt.Errorf("link param source %q has an empty field name", source)
		}
		return nil
	})
}
