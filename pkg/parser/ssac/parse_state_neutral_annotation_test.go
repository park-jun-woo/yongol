//ff:func feature=ssac-parse type=test control=iteration dimension=1
//ff:what TestHasStateNeutralComment — @state-neutral 함수 레벨 어노테이션 파싱 단위 테스트

package ssac

import (
	"testing"
)

// TestHasStateNeutralComment exercises the @state-neutral comment
// detector across four representative source shapes (present / absent /
// similar-prefix / whitespace-tolerant).
func TestHasStateNeutralComment(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want bool
	}{
		{
			name: "annotation present on its own line",
			src: `package service

// @state-neutral
// @get Workflow wf = Workflow.FindByID({ID: request.id})
// @response { ok: true }
func LikeWorkflow() {}
`,
			want: true,
		},
		{
			name: "annotation absent",
			src: `package service

// @get Workflow wf = Workflow.FindByID({ID: request.id})
// @response { ok: true }
func ExecuteWorkflow() {}
`,
			want: false,
		},
		{
			name: "similar prefix must not match (e.g. @state)",
			src: `package service

// @state workflow {Status: wf.Status} "Activate" "err"
func ActivateWorkflow() {}
`,
			want: false,
		},
		{
			name: "tolerant to leading/trailing whitespace",
			src: `package service

//    @state-neutral
func NeutralOp() {}
`,
			want: true,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			comments := parseSSaCSource(t, tc.src)
			got := hasStateNeutralComment(comments)
			if got != tc.want {
				t.Fatalf("want %v, got %v", tc.want, got)
			}
		})
	}
}
