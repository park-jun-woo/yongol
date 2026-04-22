//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what matchSensitivePattern — match a column name against sensitive-data patterns (password/token/ssn, etc.)
package ddl

import "strings"

// sensitivePatterns are column-name substrings that suggest sensitive data.
var sensitivePatterns = []string{
	// authentication credentials
	"password", "passwd", "passphrase",
	"secret", "token", "hash", "salt",
	"credential", "otp", "pin",
	// cryptographic material
	"private_key", "cipher", "encrypted",
	// financial
	"credit_card", "card_number", "cvv",
	"bank_account", "routing_number",
	// personal identifiers
	"ssn", "passport", "license_number",
	"biometric",
}

// matchSensitivePattern returns the first matching pattern, or "".
func matchSensitivePattern(colName string) string {
	lower := strings.ToLower(colName)
	for _, p := range sensitivePatterns {
		if strings.Contains(lower, p) {
			return p
		}
	}
	return ""
}
