//ff:type feature=report type=model topic=sarif
//ff:what Document — SARIF 2.1.0 top-level 문서 구조
package sarif

// Document is the top-level SARIF 2.1.0 document.
// Reference: https://docs.oasis-open.org/sarif/sarif/v2.1.0/sarif-v2.1.0.html
type Document struct {
	Schema  string `json:"$schema,omitempty"`
	Version string `json:"version"`
	Runs    []Run  `json:"runs"`
}
