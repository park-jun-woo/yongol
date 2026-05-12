//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what data-state="field=value" 비교 연산자 === 변환 검증
package stml

import "testing"

func TestResolveStateConditionEquality(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		dataVar   string
		want      string
	}{
		{
			name:      "field=value produces strict equality",
			condition: "workflow.status=draft",
			dataVar:   "data",
			want:      "data.workflow?.status === 'draft'",
		},
		{
			name:      "no equals sign passes through",
			condition: "items",
			dataVar:   "data",
			want:      "data.items",
		},
		{
			name:      "loading suffix unchanged",
			condition: "items.loading",
			dataVar:   "data",
			want:      "dataLoading",
		},
		{
			name:      "error suffix unchanged",
			condition: "items.error",
			dataVar:   "data",
			want:      "dataError",
		},
		{
			name:      "empty suffix unchanged",
			condition: "items.empty",
			dataVar:   "data",
			want:      "data.items?.length === 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveStateCondition(tt.condition, tt.dataVar)
			if got != tt.want {
				t.Errorf("resolveStateCondition(%q, %q)\n  got:  %s\n  want: %s", tt.condition, tt.dataVar, got, tt.want)
			}
		})
	}
}
