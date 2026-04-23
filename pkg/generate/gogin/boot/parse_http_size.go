//ff:func feature=gen-gogin type=util control=sequence topic=dos-guard
//ff:what parseHTTPSize — middleware.ParseSize 래퍼: 빈 문자열/에러를 ok=false 로 통일

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/gogin/middleware"

// parseHTTPSize is a wrapper over middleware.ParseSize that turns parse
// errors + empty strings into a simple ok=false signal, keeping callers
// flat. SEC validation has already surfaced any manifest-level issues so
// falling back silently is safe.
func parseHTTPSize(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}
	n, err := middleware.ParseSize(s)
	if err != nil {
		return 0, false
	}
	return n, true
}
