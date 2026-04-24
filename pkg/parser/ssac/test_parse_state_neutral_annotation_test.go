//ff:func feature=ssac-parse type=test control=iteration dimension=1
//ff:what hasStateNeutralComment — @state-neutral 함수 레벨 어노테이션 파싱 단위 테스트

package ssac

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseSSaCSource(t *testing.T, src string) []*ast.Comment {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.ssac", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	// Pick the first FuncDecl and collect preceding comments (mirrors
	// collectFuncComments) so the test exercises the same comment surface
	// the production parser sees.
	for _, d := range f.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		return collectFuncComments(f, fn.Pos())
	}
	t.Fatalf("no FuncDecl in source")
	return nil
}

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
