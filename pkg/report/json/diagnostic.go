//ff:type feature=report type=model topic=json
//ff:what Diagnostic — 개별 ERROR/WARNING 진단 (rule_id / level / file / line / col / message)
package json

// Diagnostic is a single finding (ERROR or WARNING). File / Line / Col may
// be empty / zero when the source position is unknown — consumers decide
// whether to surface that as "global" or suppress.
//
// Level is the internal Go enum value (`"ERROR"` / `"WARNING"`), NOT the
// SARIF lowercase variant. This is intentional.
type Diagnostic struct {
	RuleID  string `json:"rule_id"`
	Level   string `json:"level"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Col     int    `json:"col"`
	Message string `json:"message"`
}
