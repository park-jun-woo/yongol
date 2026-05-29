//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what pathItemVerbs 단위 테스트 (PathItem → 5개 verb/Operation 페어)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestPathItemVerbs(t *testing.T) {
	get := &openapi3.Operation{OperationID: "Get"}
	post := &openapi3.Operation{OperationID: "Post"}
	item := &openapi3.PathItem{Get: get, Post: post}

	verbs := pathItemVerbs(item)
	if len(verbs) != 5 {
		t.Fatalf("expected 5 verbs, got %d", len(verbs))
	}
	wantMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	for i, m := range wantMethods {
		if verbs[i].method != m {
			t.Errorf("verbs[%d].method = %q, want %q", i, verbs[i].method, m)
		}
	}
	if verbs[0].op != get {
		t.Errorf("GET op not wired through")
	}
	if verbs[1].op != post {
		t.Errorf("POST op not wired through")
	}
	if verbs[2].op != nil {
		t.Errorf("PUT op should be nil")
	}
}
