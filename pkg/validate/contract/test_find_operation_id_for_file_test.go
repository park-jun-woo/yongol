//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestFindOperationIDForFile — 파일명 → operationId 매핑 단위 테스트

package contract

import "testing"

func TestFindOperationIDForFile(t *testing.T) {
	cases := []struct {
		path string
		want string
	}{
		{"arts/backend/internal/service/activate_workflow.go", "ActivateWorkflow"},
		{"list_templates.go", "ListTemplates"},
		{"/tmp/a/b/c/do_something_special.go", "DoSomethingSpecial"},
		{"GetUser.go", "GetUser"},
		{"", ""},
	}
	for _, tc := range cases {
		got := findOperationIDForFile(tc.path)
		if got != tc.want {
			t.Errorf("findOperationIDForFile(%q) = %q, want %q", tc.path, got, tc.want)
		}
	}
}
