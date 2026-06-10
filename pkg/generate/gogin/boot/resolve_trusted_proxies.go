//ff:func feature=gen-gogin type=util control=sequence topic=trusted-proxy
//ff:what resolveTrustedProxies — manifest.backend.http.trusted_proxies CIDR 목록 추출 (미설정 시 nil)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveTrustedProxies returns the CIDR list declared in
// manifest.backend.http.trusted_proxies, or nil when absent (BUG-117).
//
// nil is the safe default: the generated router then calls
// r.SetTrustedProxies(nil) so gin's c.ClientIP() ignores client-supplied
// X-Forwarded-For / X-Real-IP entirely and uses RemoteAddr. Deployments
// behind a reverse proxy opt in by declaring their proxy CIDR ranges in
// the manifest (or overriding via BACKEND_HTTP_TRUSTED_PROXIES at runtime).
func resolveTrustedProxies(fs *yongol.Fullstack) []string {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.HTTP == nil {
		return nil
	}
	if len(fs.Manifest.Backend.HTTP.TrustedProxies) == 0 {
		return nil
	}
	return fs.Manifest.Backend.HTTP.TrustedProxies
}
