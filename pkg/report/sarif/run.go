//ff:type feature=report type=model topic=sarif
//ff:what Run — SARIF 도구 실행 1회분 (tool + results)
package sarif

// Run represents a single tool invocation.
type Run struct {
	Tool    Tool     `json:"tool"`
	Results []Result `json:"results"`
}
