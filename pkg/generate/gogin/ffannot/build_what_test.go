//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what test: TestBuildWhat — table-driven test verifying that a single //ff:what line is assembled correctly

package ffannot

import "testing"

func TestBuildWhat(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"  ", ""},
		{"ActivateWorkflow — 워크플로우 활성화 핸들러",
			"//ff:what ActivateWorkflow — 워크플로우 활성화 핸들러"},
		{"first\nsecond", "//ff:what first second"},
	}
	for _, tc := range cases {
		got := BuildWhat(tc.in)
		if got != tc.want {
			t.Fatalf("BuildWhat(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
