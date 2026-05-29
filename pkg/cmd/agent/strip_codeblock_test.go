//ff:func feature=agent type=test control=sequence
//ff:what TestStripCodeBlock — markdown 코드블록 울타리 제거 및 비-펜스 입력 통과 검증

package agent

import "testing"

func TestStripCodeBlock(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"no fence", "SELECT 1;", "SELECT 1;"},
		{"sql fence", "```sql\nSELECT 1;\n```", "SELECT 1;"},
		{"bare fence", "```\nfoo\n```", "foo"},
		{"surrounding ws", "  ```yaml\nkey: v\n```  ", "key: v"},
		{"fence no newline", "```", "```"},
	}
	for _, c := range cases {
		if got := stripCodeBlock(c.in); got != c.want {
			t.Errorf("%s: stripCodeBlock(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
