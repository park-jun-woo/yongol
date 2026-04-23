//ff:type feature=report type=model topic=sarif
//ff:what Rule — SARIF 규칙 메타 (id / name / shortDescription / helpUri / defaultConfiguration)
package sarif

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
