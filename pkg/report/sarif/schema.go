//ff:type feature=report type=model topic=sarif
//ff:what SARIF 2.1.0 minimal schema structs (stdlib encoding/json 만 사용)
package sarif

// Document is the top-level SARIF 2.1.0 document.
// Reference: https://docs.oasis-open.org/sarif/sarif/v2.1.0/sarif-v2.1.0.html
type Document struct {
	Schema  string `json:"$schema,omitempty"`
	Version string `json:"version"`
	Runs    []Run  `json:"runs"`
}

// Run represents a single tool invocation.
type Run struct {
	Tool    Tool     `json:"tool"`
	Results []Result `json:"results"`
}

// Tool wraps the driver descriptor.
type Tool struct {
	Driver Driver `json:"driver"`
}

// Driver identifies the analysis tool.
type Driver struct {
	Name           string `json:"name"`
	Version        string `json:"version,omitempty"`
	InformationURI string `json:"informationUri,omitempty"`
	Rules          []Rule `json:"rules,omitempty"`
}

// Rule describes an individual rule in the tool's catalog.
// PhaseF02 extended the struct beyond {id, name} to carry
// shortDescription / helpUri / defaultConfiguration / properties so that
// GitHub Code Scanning, VS Code SARIF Viewer and similar consumers can
// surface human-readable metadata for every catalogued rule.
type Rule struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name,omitempty"`
	ShortDescription     *Message               `json:"shortDescription,omitempty"`
	HelpURI              string                 `json:"helpUri,omitempty"`
	DefaultConfiguration *DefaultConfiguration  `json:"defaultConfiguration,omitempty"`
	Properties           map[string]interface{} `json:"properties,omitempty"`
}

// DefaultConfiguration carries the rule's default severity, mapped from
// the rulebook `Level` column (ERROR → "error", WARNING → "warning").
type DefaultConfiguration struct {
	Level string `json:"level,omitempty"`
}

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

// Message carries the human-readable finding text.
type Message struct {
	Text string `json:"text"`
}

// Location points at the artifact and region where a finding was detected.
type Location struct {
	PhysicalLocation PhysicalLocation `json:"physicalLocation"`
}

// PhysicalLocation binds an artifact to a region.
type PhysicalLocation struct {
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`
	Region           *Region          `json:"region,omitempty"`
}

// ArtifactLocation is the (relative) path of the source file.
type ArtifactLocation struct {
	URI string `json:"uri"`
}

// Region is a 1-based line/column span. startColumn omitted when 0.
type Region struct {
	StartLine   int `json:"startLine,omitempty"`
	StartColumn int `json:"startColumn,omitempty"`
}
