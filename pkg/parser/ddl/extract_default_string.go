//ff:func feature=manifest type=parser control=sequence
//ff:what extractDefaultString — `DEFAULT '<x>'` 패턴에서 문자열 리터럴 추출

package ddl

import "regexp"

var reDefaultString = regexp.MustCompile(`(?i)DEFAULT\s+'([^']*)'`)

// extractDefaultString returns the string literal following a `DEFAULT` clause
// in a column definition (e.g. `status VARCHAR(32) NOT NULL DEFAULT 'draft'`
// → "draft"). Returns "" when no string-literal DEFAULT present. Numeric and
// expression defaults (`DEFAULT 0`, `DEFAULT now()`) are ignored — XDM-28
// compares state machine initial states which are always string literals.
func extractDefaultString(line string) string {
	m := reDefaultString.FindStringSubmatch(line)
	if m == nil {
		return ""
	}
	return m[1]
}
