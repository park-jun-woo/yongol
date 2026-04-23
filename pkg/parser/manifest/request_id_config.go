//ff:type feature=projectconfig type=model
//ff:what RequestIDConfig — backend.error.request_id 섹션 모델

package manifest

// RequestIDConfig controls the request_id middleware. TrustUpstream=nil falls
// back to true (LB/API Gateway may already inject the header). Header=""
// falls back to "X-Request-Id".
type RequestIDConfig struct {
	TrustUpstream *bool  `yaml:"trust_upstream,omitempty"`
	Header        string `yaml:"header,omitempty"`
}
