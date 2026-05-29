//ff:func feature=gen-gogin type=test control=sequence
//ff:what collectPublicOps unit test — root security, opt-out, and public cases
package boot

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectPublicOps(t *testing.T) {
	// root security present, only Login opts out
	doc := buildDoc(
		[]opSpec{
			{path: "/login", method: "POST", opID: "Login", sec: &openapi3.SecurityRequirements{}},
			{path: "/workflows", method: "POST", opID: "CreateWorkflow", sec: nil},
			{path: "/workflows", method: "GET", opID: "ListWorkflows", sec: nil},
		},
		true,
	)
	if got := collectPublicOps(doc); !equalStrings(got, []string{"Login"}) {
		t.Errorf("case1: got %v want [Login]", got)
	}

	// no root security → all ops are public
	doc2 := buildDoc(
		[]opSpec{
			{path: "/a", method: "GET", opID: "A", sec: nil},
			{path: "/b", method: "POST", opID: "B", sec: nil},
		},
		false,
	)
	if got := collectPublicOps(doc2); !equalStrings(got, []string{"A", "B"}) {
		t.Errorf("case2: got %v want [A B]", got)
	}

	// all ops protected
	doc3 := buildDoc(
		[]opSpec{
			{path: "/x", method: "POST", opID: "X", sec: nil},
			{path: "/y", method: "POST", opID: "Y", sec: nil},
		},
		true,
	)
	if got := collectPublicOps(doc3); len(got) != 0 {
		t.Errorf("case3: got %v want empty", got)
	}

	// nil doc
	if got := collectPublicOps(nil); got != nil {
		t.Errorf("case4: got %v want nil", got)
	}
}
