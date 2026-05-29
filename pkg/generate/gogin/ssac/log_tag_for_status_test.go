//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what logTagForStatus 단위 테스트 (4xx→"4xx", 그 외→"5xx")

package ssac

import "testing"

func TestLogTagForStatus(t *testing.T) {
	cases := map[int]string{
		400: "4xx",
		404: "4xx",
		499: "4xx",
		500: "5xx",
		200: "5xx",
	}
	for in, want := range cases {
		if got := logTagForStatus(in); got != want {
			t.Errorf("logTagForStatus(%d) = %q, want %q", in, got, want)
		}
	}
}
