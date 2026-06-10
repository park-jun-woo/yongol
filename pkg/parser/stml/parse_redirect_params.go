//ff:func feature=stml-parse type=parser control=sequence
//ff:what ParseRedirectParams — data-redirect-params 값("respField -> Segment, ...")을 LinkParamBind 목록으로 파싱 (TM-33 구문 근거)
package stml

import (
	"fmt"
	"strings"
)

// ParseRedirectParams parses a data-redirect-params attribute value of the
// form "<source> -> <SegmentName>[, ...]" into LinkParamBind entries
// (page-flow Phase008) — the same value-based "a -> b" grammar as
// data-link-params (parseParamBinds core). Sources are unprefixed 2xx
// response fields of the action's operation (the only data in scope right
// after the action succeeds — the same namespace tier as the data-capture
// left-hand side) or "route.<Name>" (forwarding a current-page param).
// item.<Field> rows are not in scope after an action and are rejected.
// The "-> <SegmentName>" part may be elided; TM-33 then requires the
// target route to have exactly one required segment. Any format violation
// returns an error; TM-33 re-parses the raw attribute at validate time to
// surface it as a diagnostic (the ParseCapture / TM-20 split).
func ParseRedirectParams(raw string) ([]LinkParamBind, error) {
	return parseParamBinds(raw, "redirect param", func(source string) error {
		if source == "" {
			return fmt.Errorf("redirect param binding has an empty source; expected \"<respField>\" or \"route.<Name>\"")
		}
		if strings.HasPrefix(source, "item.") {
			return fmt.Errorf("redirect param source %q is invalid: no row item is in scope after an action — use a 2xx response field or \"route.<Name>\"", source)
		}
		if source == "route." {
			return fmt.Errorf("redirect param source %q has an empty field name", source)
		}
		return nil
	})
}
