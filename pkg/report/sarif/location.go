//ff:type feature=report type=model topic=sarif
//ff:what Location — SARIF 위치 (physicalLocation 래퍼)
package sarif

// Location points at the artifact and region where a finding was detected.
type Location struct {
	PhysicalLocation PhysicalLocation `json:"physicalLocation"`
}
