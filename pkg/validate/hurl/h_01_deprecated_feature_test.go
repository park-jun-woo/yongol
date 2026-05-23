//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-structural
//ff:what h01DeprecatedFeature — deprecated .feature 파일 탐지 검증

package hurl

import (
	"testing"
)

func TestH01DeprecatedFeature(t *testing.T) {
	cases := []TestH01DeprecatedFeatureCase{
		{
			name:      "no_feature_files_no_diag",
			files:     map[string]string{"tests/login.hurl": ""},
			wantCount: 0,
		},
		{
			name:      "feature_in_tests_produces_diag",
			files:     map[string]string{"tests/login.feature": ""},
			wantCount: 1,
		},
		{
			name:      "feature_in_scenario_produces_diag",
			files:     map[string]string{"scenario/signup.feature": ""},
			wantCount: 1,
		},
		{
			name: "features_in_both_dirs",
			files: map[string]string{
				"tests/login.feature":     "",
				"scenario/signup.feature": "",
			},
			wantCount: 2,
		},
		{
			name: "multiple_features_in_tests",
			files: map[string]string{
				"tests/login.feature":  "",
				"tests/signup.feature": "",
			},
			wantCount: 2,
		},
		{
			name:      "empty_dirs_no_diag",
			files:     nil,
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runH01DeprecatedFeature(t, c)
		})
	}
}
