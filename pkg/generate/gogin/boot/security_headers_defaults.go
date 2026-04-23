package boot

// Default header values applied when the manifest omits the security_headers
// block or an individual field. production profile enables every header.

const (
	defaultSHProfile       = "production"
	defaultHSTSMaxAge      = 31536000 // 1 year
	defaultHSTSIncludeSubs = true
	defaultHSTSPreload     = false
	defaultXFrameOptions   = "DENY"
	defaultReferrerPolicy  = "strict-origin-when-cross-origin"
)
