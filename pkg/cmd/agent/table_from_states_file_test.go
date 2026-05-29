//ff:func feature=agent type=test control=sequence
//ff:what TestTableFromStatesFile — states 파일명/경로에서 테이블명(.md 제거) 추출 검증

package agent

import "testing"

func TestTableFromStatesFile(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"states/order.md", "order"},
		{"order.md", "order"},
		{"states/workflow.md", "workflow"},
		{"plain", "plain"},
	}
	for _, c := range cases {
		if got := tableFromStatesFile(c.in); got != c.want {
			t.Errorf("tableFromStatesFile(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
