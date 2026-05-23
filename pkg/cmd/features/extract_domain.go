//ff:func feature=features type=helper control=sequence
//ff:what extractDomain — HTTP path 에서 도메인명 추출 (e.g. "POST /workflows/{id}" → "workflow")

package features

import "strings"

// extractDomain extracts the domain name from an HTTP path like "POST /workflows/{id}".
func extractDomain(httpPath string) string {
	parts := strings.Fields(httpPath)
	if len(parts) < 2 {
		return "unknown"
	}
	uri := parts[1]
	uri = strings.TrimPrefix(uri, "/")
	seg := strings.SplitN(uri, "/", 2)[0]
	if seg == "" {
		return "unknown"
	}
	if idx := strings.Index(seg, "-"); idx > 0 {
		seg = seg[:idx]
	}
	seg = strings.TrimSuffix(seg, "s")
	if seg == "" {
		return "unknown"
	}
	return seg
}
