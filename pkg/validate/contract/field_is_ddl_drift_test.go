//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestFieldIsDDLDrift — selector 파싱 및 Field 부분 대조 단위 테스트

package contract

import "testing"

func TestFieldIsDDLDrift(t *testing.T) {
	expected := map[string]bool{"email": true}
	cases := []struct {
		selector string
		drift    bool
	}{
		{"u.Email", false},
		{"user.Email", false},
		{"u.DeletedAt", true},
		{"NoDot", false},
		{"trailing.", false},
	}
	for _, tc := range cases {
		got := fieldIsDDLDrift(tc.selector, expected)
		if got != tc.drift {
			t.Errorf("fieldIsDDLDrift(%q) = %v, want %v", tc.selector, got, tc.drift)
		}
	}
}
