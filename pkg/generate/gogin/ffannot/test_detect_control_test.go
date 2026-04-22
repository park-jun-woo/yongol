//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what test: TestDetectControl — body 라인 제어구조 추론 검증

package ffannot

import "testing"

func TestDetectControl(t *testing.T) {
	cases := []struct {
		name string
		body []string
		want string
	}{
		{
			"plain-sequence",
			[]string{
				`x := 1`,
				`return x`,
			},
			ControlSequence,
		},
		{
			"loop-detected",
			[]string{
				`for i, row := range rows {`,
				`    result[i] = row`,
				`}`,
				`return result`,
			},
			ControlIteration,
		},
		{
			"switch-detected",
			[]string{
				`switch kind {`,
				`case "a":`,
				`    return 1`,
				`}`,
				`return 0`,
			},
			ControlSelection,
		},
		{
			"nested-switch-inside-loop-is-iteration",
			[]string{
				`for _, x := range xs {`,
				`    switch x {`,
				`    case 1:`,
				`        return 1`,
				`    }`,
				`}`,
			},
			ControlIteration,
		},
		{
			"switch-in-comment-ignored",
			[]string{
				`// switch case discussion`,
				`return 0`,
			},
			ControlSequence,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := DetectControl(tc.body)
			if got != tc.want {
				t.Fatalf("DetectControl() = %q, want %q", got, tc.want)
			}
		})
	}
}
