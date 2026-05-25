//ff:type feature=projectconfig type=model topic=rate-limit
//ff:what RateLimitEntry — 단일 operationId 의 rate/period/key 설정

package manifest

// RateLimitEntry represents a single rate limit rule for one operationId.
type RateLimitEntry struct {
	Rate   int    `yaml:"rate"`
	Period string `yaml:"period"`
	Key    string `yaml:"key"`
}
