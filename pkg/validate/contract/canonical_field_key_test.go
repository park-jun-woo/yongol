//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestCanonicalFieldKey — DDL/Go 필드 이름을 소문자·무언더스코어 키로 정규화 검증

package contract

import "testing"

func TestCanonicalFieldKey(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"pascal initialism", "OrgID", "orgid"},
		{"camel", "orgId", "orgid"},
		{"snake", "org_id", "orgid"},
		{"screaming snake", "ORG_ID", "orgid"},
		{"plain", "email", "email"},
		{"empty", "", ""},
		{"multi underscore", "created_at_utc", "createdatutc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canonicalFieldKey(tt.in); got != tt.want {
				t.Fatalf("canonicalFieldKey(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
