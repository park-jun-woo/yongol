//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what renderRemoveQueriesExpr — 빈 목록/단일 op/다중 op join 의 removeQueries 표현식 검증

package stml

import "testing"

func TestRenderRemoveQueriesExpr(t *testing.T) {
	tests := []struct {
		name      string
		removeOps []string
		want      string
	}{
		{
			name:      "empty list yields empty string",
			removeOps: nil,
			want:      "",
		},
		{
			name:      "single op",
			removeOps: []string{"GetRoom"},
			want:      "queryClient.removeQueries({ queryKey: ['GetRoom'] })",
		},
		{
			name:      "multiple ops joined with newline indent",
			removeOps: []string{"GetRoom", "GetBuilding"},
			want: "queryClient.removeQueries({ queryKey: ['GetRoom'] })" +
				"\n      queryClient.removeQueries({ queryKey: ['GetBuilding'] })",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderRemoveQueriesExpr(tt.removeOps); got != tt.want {
				t.Errorf("renderRemoveQueriesExpr() = %q, want %q", got, tt.want)
			}
		})
	}
}
