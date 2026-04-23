//ff:func feature=rule type=test control=iteration dimension=1
//ff:what TestBaseSpec_Validate — BaseSpec.Validate 가 Rule/Level 빈 문자열을 거부하는지 검증

package rule

import (
	"testing"
)

func TestBaseSpec_Validate(t *testing.T) {
	cases := []struct {
		name    string
		spec    BaseSpec
		wantErr bool
		wantSub string
	}{
		{
			name:    "ok",
			spec:    BaseSpec{Rule: "TEST-1", Level: "ERROR", Message: "hello"},
			wantErr: false,
		},
		{
			name:    "empty-rule",
			spec:    BaseSpec{Level: "ERROR", Message: "m"},
			wantErr: true,
			wantSub: "Rule",
		},
		{
			name:    "empty-level",
			spec:    BaseSpec{Rule: "TEST-1", Message: "m"},
			wantErr: true,
			wantSub: "Level",
		},
		{
			name:    "both-empty",
			spec:    BaseSpec{},
			wantErr: true,
			wantSub: "Rule",
		},
		{
			name:    "message-empty-ok",
			spec:    BaseSpec{Rule: "X", Level: "WARNING"},
			wantErr: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, runBaseSpecValidateCase(tc.spec, tc.wantErr, tc.wantSub))
	}
}
