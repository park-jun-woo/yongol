//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what renderEnabledWhenExpr이 data-enabled-when 가드를 JSX boolean 식으로 변환하는지 검증
package stml

import "testing"

func TestRenderEnabledWhenExpr(t *testing.T) {
	tests := []struct {
		name        string
		enabledWhen string
		dataVar     string
		want        string
	}{
		{
			name:        "empty condition returns empty",
			enabledWhen: "",
			dataVar:     "data",
			want:        "",
		},
		{
			name:        "equality guard converts to strict equality",
			enabledWhen: "workflow.status=draft",
			dataVar:     "data",
			want:        "data.workflow?.status === 'draft'",
		},
		{
			name:        "plain truthy guard passes through",
			enabledWhen: "items",
			dataVar:     "data",
			want:        "data.items",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderEnabledWhenExpr(tt.enabledWhen, tt.dataVar)
			if got != tt.want {
				t.Errorf("renderEnabledWhenExpr(%q, %q)\n  got:  %s\n  want: %s", tt.enabledWhen, tt.dataVar, got, tt.want)
			}
		})
	}
}
