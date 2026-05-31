//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestMatchSensitivePattern(t *testing.T) {
	tests := []struct {
		col  string
		want string
	}{
		{"user_password", "password"},
		{"PassWord", "password"},
		{"api_token", "token"},
		{"ssn", "ssn"},
		{"credit_card_no", "credit_card"},
		{"private_key", "private_key"},
		{"username", ""},
		{"email", ""},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.col, func(t *testing.T) {
			if got := matchSensitivePattern(tt.col); got != tt.want {
				t.Errorf("matchSensitivePattern(%q) = %q, want %q", tt.col, got, tt.want)
			}
		})
	}
}
