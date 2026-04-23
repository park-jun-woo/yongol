//ff:type feature=report type=model topic=sarif
//ff:what Tool — SARIF driver 래퍼
package sarif

// Tool wraps the driver descriptor.
type Tool struct {
	Driver Driver `json:"driver"`
}
