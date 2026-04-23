//ff:type feature=report type=model topic=sarif
//ff:what ArtifactLocation — (상대) 소스 파일 경로
package sarif

// ArtifactLocation is the (relative) path of the source file.
type ArtifactLocation struct {
	URI string `json:"uri"`
}
