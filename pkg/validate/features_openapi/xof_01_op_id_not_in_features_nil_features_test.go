//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what XOF-01 — Features nil 시 단락 테스트

package features_openapi

import (
	"testing"
)

func TestXOF01_OpIDNotInFeatures_NilFeatures(t *testing.T) {
	// When features is nil the rule short-circuits and returns nil (same as NilGround).
	fs := buildFSForXOF01([]string{"CreateWorkflow"}, nil, nil)
	diags := xof01OpIDNotInFeatures(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags with nil features, got %d", len(diags))
	}
}
