//ff:type feature=report type=model topic=sarif
//ff:what Driver — SARIF 도구 식별자 (name / version / informationUri / rules)
package sarif

// Driver identifies the analysis tool.
type Driver struct {
	Name           string `json:"name"`
	Version        string `json:"version,omitempty"`
	InformationURI string `json:"informationUri,omitempty"`
	Rules          []Rule `json:"rules,omitempty"`
}
