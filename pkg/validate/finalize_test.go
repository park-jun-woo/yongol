//ff:func feature=validate type=test control=iteration dimension=1
//ff:what finalize — ERROR 유무에 따라 StatusPass/StatusFail 분기 검증

package validate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestFinalize(t *testing.T) {
	cases := []struct {
		name       string
		stepName   string
		diags      []diagnostic.Diagnostic
		wantStatus Status
	}{
		{
			name:       "no_diagnostics_returns_pass",
			stepName:   "step-empty",
			diags:      nil,
			wantStatus: StatusPass,
		},
		{
			name:     "only_warnings_returns_pass",
			stepName: "step-warn",
			diags: []diagnostic.Diagnostic{
				{Level: diagnostic.LevelWarning, Message: "minor issue"},
			},
			wantStatus: StatusPass,
		},
		{
			name:     "single_error_returns_fail",
			stepName: "step-err",
			diags: []diagnostic.Diagnostic{
				{Level: diagnostic.LevelError, Message: "broken"},
			},
			wantStatus: StatusFail,
		},
		{
			name:     "error_among_warnings_returns_fail",
			stepName: "step-mixed",
			diags: []diagnostic.Diagnostic{
				{Level: diagnostic.LevelWarning, Message: "w1"},
				{Level: diagnostic.LevelError, Message: "e1"},
				{Level: diagnostic.LevelWarning, Message: "w2"},
			},
			wantStatus: StatusFail,
		},
		{
			name:       "empty_slice_returns_pass",
			stepName:   "step-empty-slice",
			diags:      []diagnostic.Diagnostic{},
			wantStatus: StatusPass,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runFinalizeCase(t, c.stepName, c.diags, c.wantStatus)
		})
	}
}
