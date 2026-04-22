//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what matchSensitivePattern — 컬럼 이름에서 민감 데이터 패턴 매칭 (password/token/ssn 등)
package ddl

import "strings"

// sensitivePatterns are column-name substrings that suggest sensitive data.
var sensitivePatterns = []string{
	// 인증 정보
	"password", "passwd", "passphrase",
	"secret", "token", "hash", "salt",
	"credential", "otp", "pin",
	// 암호화
	"private_key", "cipher", "encrypted",
	// 금융
	"credit_card", "card_number", "cvv",
	"bank_account", "routing_number",
	// 개인식별
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
