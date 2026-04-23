//ff:type feature=projectconfig type=model
//ff:what HTTPOverride — backend.http.overrides 엔트리 모델 (operationId 별 limit 오버라이드)

package manifest

// HTTPOverride is a per-operationId override. Either or both fields may be
// set; unset fields fall back to the global limits.
type HTTPOverride struct {
	BodyLimit      string `yaml:"body_limit"`
	MultipartLimit string `yaml:"multipart_limit"`
}
