//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what XFO-01 — Features nil 시 단락 테스트

package features_openapi

import (
	"testing"
)

func TestXFO01_OpNotInOpenAPI_NilFeatures(t *testing.T) {
	fs := buildFSForXFO01([]string{"CreateWorkflow"}, nil)
	diags := xfo01OpNotInOpenAPI(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags with nil features, got %d", len(diags))
	}
}
