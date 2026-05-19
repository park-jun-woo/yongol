//ff:func feature=cli-init type=util control=sequence
//ff:what extractDomain — extracts domain name from an HTTP path for SSaC directory grouping

package cliinit

import "strings"

// extractDomain extracts the domain name from an HTTP path like "POST /workflows/{id}".
// It takes the first path segment and singularises simple plurals by stripping
// trailing "s". Hyphens are collapsed: "/audit-logs" → "audit".
func extractDomain(httpPath string) string {
	// Split "POST /workflows/{id}" → ["POST", "/workflows/{id}"]
	parts := strings.Fields(httpPath)
	if len(parts) < 2 {
		return "unknown"
	}
	uri := parts[1]
	// Trim leading slash and take first segment.
	uri = strings.TrimPrefix(uri, "/")
	seg := strings.SplitN(uri, "/", 2)[0]
	if seg == "" {
		return "unknown"
	}
	// Handle hyphenated names: "audit-logs" → take first part "audit".
	if idx := strings.Index(seg, "-"); idx > 0 {
		seg = seg[:idx]
	}
	// Simple deplural: strip trailing "s".
	seg = strings.TrimSuffix(seg, "s")
	if seg == "" {
		return "unknown"
	}
	return seg
}
