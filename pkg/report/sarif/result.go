//ff:type feature=report type=model topic=sarif
//ff:what Result — 단일 진단 (ruleId / ruleIndex / level / message / locations)
package sarif

// Result is a single diagnostic finding.
// RuleIndex (0-based) points into runs[0].tool.driver.rules[] for consumers
// that resolve rule metadata by index rather than string id.
type Result struct {
	RuleID    string     `json:"ruleId,omitempty"`
	RuleIndex *int       `json:"ruleIndex,omitempty"`
	Level     string     `json:"level,omitempty"`
	Message   Message    `json:"message"`
	Locations []Location `json:"locations,omitempty"`
}
