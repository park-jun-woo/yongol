//ff:func feature=report type=util control=selection topic=sarif
//ff:what mapLevel — diagnostic.Level → SARIF level 문자열 ("error"/"warning"/"note")
package sarif

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// mapLevel converts our internal severity into SARIF's level enum.
func mapLevel(l diagnostic.Level) string {
	switch l {
	case diagnostic.LevelError:
		return "error"
	case diagnostic.LevelWarning:
		return "warning"
	}
	return "note"
}
