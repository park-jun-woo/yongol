//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what 가드 분기 codegen — &&/||/!/괄호/비교 연산자 AST→JSX 변환 검증

package stml

import "testing"

func TestResolveStateConditionGuard(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		want      string
	}{
		{
			name:      "and combination",
			condition: "workflow.status=active && currentUser.Role=owner",
			want:      "data.workflow?.status === 'active' && data.currentUser?.Role === 'owner'",
		},
		{
			name:      "or combination",
			condition: "workflow.status=active || workflow.status=draft",
			want:      "data.workflow?.status === 'active' || data.workflow?.status === 'draft'",
		},
		{
			name:      "negation",
			condition: "!workflow.status=draft",
			want:      "!data.workflow?.status === 'draft'",
		},
		{
			name:      "group",
			condition: "(workflow.status=active || workflow.status=draft)",
			want:      "(data.workflow?.status === 'active' || data.workflow?.status === 'draft')",
		},
		{
			name:      "not equal operator",
			condition: "workflow.status!=draft && order.count>=3",
			want:      "data.workflow?.status !== 'draft' && data.order?.count >= '3'",
		},
		{
			name:      "lifecycle inside guard",
			condition: "!items.list.loading && items.list.empty",
			want:      "!dataLoading && data.items?.list?.length === 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveStateCondition(tt.condition, "data")
			if got != tt.want {
				t.Errorf("resolveStateCondition(%q)\n  got:  %s\n  want: %s", tt.condition, got, tt.want)
			}
		})
	}
}
