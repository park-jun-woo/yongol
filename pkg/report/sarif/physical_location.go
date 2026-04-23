//ff:type feature=report type=model topic=sarif
//ff:what PhysicalLocation — artifactLocation + region 바인딩
package sarif

// PhysicalLocation binds an artifact to a region.
type PhysicalLocation struct {
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`
	Region           *Region          `json:"region,omitempty"`
}
