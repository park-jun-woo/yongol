//ff:func feature=gen-gogin type=util control=sequence topic=request-id
//ff:what resolveRequestIDConfig — trust_upstream + header 기본값 (true, "X-Request-Id")

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveRequestIDConfig reads manifest.backend.error.request_id with
// sensible defaults when the block is absent. trust_upstream defaults to
// true; header defaults to "X-Request-Id".
func resolveRequestIDConfig(fs *yongol.Fullstack) (bool, string) {
	trust := true
	header := "X-Request-Id"
	if fs == nil || fs.Manifest == nil {
		return trust, header
	}
	e := fs.Manifest.Backend.Error
	if e == nil || e.RequestID == nil {
		return trust, header
	}
	if e.RequestID.TrustUpstream != nil {
		trust = *e.RequestID.TrustUpstream
	}
	if e.RequestID.Header != "" {
		header = e.RequestID.Header
	}
	return trust, header
}
