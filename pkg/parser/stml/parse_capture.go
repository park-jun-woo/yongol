//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what ParseCapture — data-capture 값("a -> b, c -> d")을 CaptureBind 목록으로 파싱 (TM-20 구문 근거)
package stml

import (
	"fmt"
	"strings"
)

// ParseCapture parses a data-capture attribute value of the form
// "<respField> -> <sink>[, <respField> -> <sink>...]" into CaptureBind
// entries. The sink namespace is restricted to "auth.token",
// "auth.refresh" and "auth.claims.<name>" — a claim captured from the
// login response body for the sitemap data-roles menu filter
// (plans/stml/sitemap Phase005) — ("session.*" collides with the SSaC
// built-in session package and is rejected). Any format violation returns
// an error; TM-20 re-parses the raw attribute at validate time to surface
// that error as a diagnostic, mirroring the ParseGuard / TM-17 split.
func ParseCapture(raw string) ([]CaptureBind, error) {
	var out []CaptureBind
	for _, seg := range strings.Split(raw, ",") {
		seg = strings.TrimSpace(seg)
		if seg == "" {
			return nil, fmt.Errorf("empty capture binding; expected \"<respField> -> <sink>\"")
		}
		parts := strings.Split(seg, "->")
		if len(parts) != 2 {
			return nil, fmt.Errorf("capture binding %q must be \"<respField> -> <sink>\"", seg)
		}
		field := strings.TrimSpace(parts[0])
		sink := strings.TrimSpace(parts[1])
		if field == "" {
			return nil, fmt.Errorf("capture binding %q has an empty response field", seg)
		}
		if _, claims := ClaimsSinkName(sink); sink != "auth.token" && sink != "auth.refresh" && !claims {
			return nil, fmt.Errorf("capture sink %q is not allowed; use \"auth.token\", \"auth.refresh\" or \"auth.claims.<name>\"", sink)
		}
		out = append(out, CaptureBind{RespField: field, Sink: sink})
	}
	return out, nil
}
