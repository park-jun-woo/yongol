//ff:func feature=rule type=test control=iteration dimension=1
//ff:what 확장된 goInitialisms 사전이 yongol 스택 약자 위반을 감지하는지 검증

package rule

import "testing"

// TestViolatesInitialism_Extended covers the entries added in PhaseV11
// (JWT/SQLC/CORS/CSRF/OAUTH/GORM/OIDC/SAML/JWK/JWS/MD5/SHA1/SHA256/
// CDN/QR/OTP/TOTP/HOTP) and regression-guards existing entries
// (OrgId/UrlBox/HtmlTag).
func TestViolatesInitialism_Extended(t *testing.T) {
	cases := []struct {
		in          string
		wantCorrect string
		wantViol    bool
	}{
		// --- New PhaseV11 entries (positive: Pascal-cased violation detected) ---
		{"JwtToken", "JWTToken", true},
		{"SqlcQuery", "SQLCQuery", true},
		{"CorsConfig", "CORSConfig", true},
		{"CsrfToken", "CSRFToken", true},
		{"OauthClient", "OAUTHClient", true},
		{"GormDB", "GORMDB", true},
		{"OidcProvider", "OIDCProvider", true},
		{"SamlAssertion", "SAMLAssertion", true},
		{"JwkSet", "JWKSet", true},
		{"JwsHeader", "JWSHeader", true},
		{"Md5Hash", "MD5Hash", true},
		{"Sha1Sum", "SHA1Sum", true},
		{"Sha256Sum", "SHA256Sum", true},
		{"CdnHost", "CDNHost", true},
		{"QrCode", "QRCode", true},
		{"OtpSecret", "OTPSecret", true},
		{"TotpCode", "TOTPCode", true},
		{"HotpCode", "HOTPCode", true},

		// --- Already-correct forms: no violation ---
		{"JWTToken", "", false},
		{"SQLCQuery", "", false},
		{"CORSConfig", "", false},
		{"MD5Hash", "", false},
		{"SHA256Sum", "", false},

		// --- Regression: existing entries still detected ---
		{"OrgId", "OrgID", true},
		{"UrlBox", "URLBox", true},
		{"HtmlTag", "HTMLTag", true},

		// --- Unrelated identifiers: no violation ---
		{"Email", "", false},
		{"ID", "", false},
		{"", "", false},
	}
	for _, tc := range cases {
		got, violated := ViolatesInitialism(tc.in)
		if violated != tc.wantViol {
			t.Errorf("ViolatesInitialism(%q) violation = %v; want %v", tc.in, violated, tc.wantViol)
			continue
		}
		if got != tc.wantCorrect {
			t.Errorf("ViolatesInitialism(%q) correct = %q; want %q", tc.in, got, tc.wantCorrect)
		}
	}
}
