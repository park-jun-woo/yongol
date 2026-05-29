//ff:func feature=cli-init type=test control=iteration dimension=1
//ff:what TestValidateProjectID — ValidateProjectID accepts well-formed ids and rejects malformed ones

package cliinit

import "testing"

func TestValidateProjectID(t *testing.T) {
	cases := []struct {
		id      string
		wantErr bool
	}{
		{"Zenflow", false},
		{"zen_flow", false},
		{"MyApp", false},
		{"a", false},
		{"A1", false},
		{"my_app2", false},
		{"", true},
		{"1abc", true},
		{"my-app", true},
		{"my.app", true},
		{"my/app", true},
		{"my app", true},
	}
	for _, tc := range cases {
		err := ValidateProjectID(tc.id)
		if tc.wantErr && err == nil {
			t.Errorf("ValidateProjectID(%q) want error, got nil", tc.id)
		}
		if !tc.wantErr && err != nil {
			t.Errorf("ValidateProjectID(%q) unexpected error: %v", tc.id, err)
		}
	}
}
