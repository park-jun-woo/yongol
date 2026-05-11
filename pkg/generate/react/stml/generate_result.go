//ff:type feature=stml-gen type=model
//ff:what 코드 생성 결과 메타데이터를 담는 구조체
package stml

// GenerateResult contains generation output metadata.
type GenerateResult struct {
	Pages        int
	Dependencies map[string]string // package name → version range
}
