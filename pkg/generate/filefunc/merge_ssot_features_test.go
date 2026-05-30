//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestMergeSSOTFeatures — SSOT feature 병합 + 설명 우선순위 해소 검증

package filefunc

import "testing"

func TestMergeSSOTFeatures(t *testing.T) {
	dst := map[string]string{}
	ssot := map[string]string{
		"workflow": "",            // empty -> infra baseline
		"auth":     "custom auth", // non-empty SSOT desc kept
		"unknownx": "",            // empty + no infra -> fallback
	}
	mergeSSOTFeatures(dst, ssot)

	if got := dst["workflow"]; got != "workflow execution, cloning, and run scenarios" {
		t.Errorf("workflow: expected infra baseline, got %q", got)
	}
	if got := dst["auth"]; got != "custom auth" {
		t.Errorf("auth: expected SSOT desc, got %q", got)
	}
	if got := dst["unknownx"]; got != fallbackDescription {
		t.Errorf("unknownx: expected fallback, got %q", got)
	}
}
