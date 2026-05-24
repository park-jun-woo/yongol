//ff:func feature=validate type=test control=iteration dimension=1 topic=policy-check
//ff:what ownerLookupName — snake_case/단일 단어/빈 문자열 PascalCase 변환 검증

package query_rego

import "testing"

func TestOwnerLookupName(t *testing.T) {
	cases := []struct {
		name     string
		resource string
		want     string
	}{
		{name: "single_word", resource: "order", want: "OwnerLookupOrder"},
		{name: "snake_case", resource: "execution_log", want: "OwnerLookupExecutionLog"},
		{name: "triple_snake", resource: "org_member_role", want: "OwnerLookupOrgMemberRole"},
		{name: "already_pascal_like", resource: "Workflow", want: "OwnerLookupWorkflow"},
		{name: "empty", resource: "", want: "OwnerLookup"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ownerLookupName(c.resource)
			if got != c.want {
				t.Errorf("ownerLookupName(%q) = %q, want %q", c.resource, got, c.want)
			}
		})
	}
}
