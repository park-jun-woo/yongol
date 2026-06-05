//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what mergeDisabledExpr이 data-enabled-when을 base에 || !(...)로 병합하는지 검증
package stml

import "testing"

func TestMergeDisabledExpr(t *testing.T) {
	tests := []struct {
		name        string
		base        string
		enabledWhen string
		dataVar     string
		want        string
	}{
		{
			name:        "empty enabledWhen returns base unchanged",
			base:        "mutation.isPending",
			enabledWhen: "",
			dataVar:     "data",
			want:        "mutation.isPending",
		},
		{
			name:        "equality guard OR-merged as negation",
			base:        "mutation.isPending",
			enabledWhen: "workflow.status=draft",
			dataVar:     "data",
			want:        "mutation.isPending || !(data.workflow?.status === 'draft')",
		},
		{
			name:        "passthrough guard wrapped in negation",
			base:        "m.isPending",
			enabledWhen: "items",
			dataVar:     "data",
			want:        "m.isPending || !(data.items)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeDisabledExpr(tt.base, tt.enabledWhen, tt.dataVar)
			if got != tt.want {
				t.Errorf("mergeDisabledExpr(%q, %q, %q)\n  got:  %s\n  want: %s", tt.base, tt.enabledWhen, tt.dataVar, got, tt.want)
			}
		})
	}
}
