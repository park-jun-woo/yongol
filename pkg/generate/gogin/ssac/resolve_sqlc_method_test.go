//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what resolveSQLCMethod 단위 테스트 ("Model.Method" → "ModelMethod")

package ssac

import "testing"

func TestResolveSQLCMethod(t *testing.T) {
	cases := map[string]string{
		"Workflow.FindByID":   "WorkflowFindByID",
		"auth.VerifyPassword": "authVerifyPassword",
		"User.Create":         "UserCreate",
		"NoDot":               "NoDot",
		"a.b.c":               "abc",
	}
	for in, want := range cases {
		if got := resolveSQLCMethod(in); got != want {
			t.Errorf("resolveSQLCMethod(%q) = %q, want %q", in, got, want)
		}
	}
}
