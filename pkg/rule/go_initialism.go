//ff:func feature=rule type=util control=iteration dimension=1
//ff:what Go initialism dictionary + ViolatesInitialism — checks whether a PascalCase identifier violates Go naming convention
package rule

import "strings"

// goInitialisms is an extended set that adds yongol stack acronyms to the
// common initialisms from the golint convention. Go naming convention: all
// acronyms must be uppercase (Id → ID, Url → URL, Jwt → JWT, …). Sorted alphabetically.
var goInitialisms = map[string]bool{
	"ACL":    true,
	"API":    true,
	"ASCII":  true,
	"CDN":    true,
	"CORS":   true,
	"CPU":    true,
	"CSRF":   true,
	"CSS":    true,
	"DNS":    true,
	"EOF":    true,
	"GORM":   true,
	"GUID":   true,
	"HOTP":   true,
	"HTML":   true,
	"HTTP":   true,
	"HTTPS":  true,
	"ID":     true,
	"IP":     true,
	"JSON":   true,
	"JWK":    true,
	"JWS":    true,
	"JWT":    true,
	"LHS":    true,
	"MD5":    true,
	"OAUTH":  true,
	"OIDC":   true,
	"OK":     true,
	"OTP":    true,
	"QPS":    true,
	"QR":     true,
	"RAM":    true,
	"RHS":    true,
	"RPC":    true,
	"SAML":   true,
	"SHA1":   true,
	"SHA256": true,
	"SLA":    true,
	"SMTP":   true,
	"SQL":    true,
	"SQLC":   true,
	"SSH":    true,
	"TCP":    true,
	"TLS":    true,
	"TOTP":   true,
	"TTL":    true,
	"UDP":    true,
	"UI":     true,
	"UID":    true,
	"URI":    true,
	"URL":    true,
	"UTF8":   true,
	"UUID":   true,
	"VM":     true,
	"XML":    true,
	"XMPP":   true,
	"XSRF":   true,
	"XSS":    true,
}

// ViolatesInitialism reports whether a PascalCase identifier violates Go
// initialism conventions. Examples:
//   - "Id"     → ("ID", true)
//   - "OrgId"  → ("OrgID", true)
//   - "UrlBox" → ("URLBox", true)
//   - "Email"  → ("", false)
//   - "ID"     → ("", false) — already correct
//
// Returns (correctForm, true) when violation detected; (s, false) otherwise.
func ViolatesInitialism(s string) (string, bool) {
	if s == "" {
		return "", false
	}
	// tokenize PascalCase: each uppercase letter starts a new token.
	tokens := splitPascal(s)
	fixed := make([]string, len(tokens))
	violated := false
	for i, tok := range tokens {
		up := strings.ToUpper(tok)
		if goInitialisms[up] && tok != up {
			fixed[i] = up
			violated = true
		} else {
			fixed[i] = tok
		}
	}
	if !violated {
		return "", false
	}
	return strings.Join(fixed, ""), true
}
