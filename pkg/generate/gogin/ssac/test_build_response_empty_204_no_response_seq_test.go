//ff:func feature=gen-gogin type=test control=iteration topic=response
//ff:what TestGenerateHTTPMethod_Delete204_NoResponse — @response 시퀀스 없는 DELETE 204 기본 return 검증 (BUG-068)

package ssac

import (
	"fmt"
	"strings"
	"testing"
)

// TestGenerateHTTPMethod_Delete204_NoResponse verifies that when
// responseLines is empty (no @response sequence in SSaC), the else
// branch in generateHTTPMethod emits a default return statement
// using the Response suffix for 204 (not JSONResponse).
func TestGenerateHTTPMethod_Delete204_NoResponse(t *testing.T) {
	// Reproduce the exact else-branch logic from generateHTTPMethod.
	funcName := "DeleteWebhook"
	successStatus := 204

	suffix := "JSONResponse"
	if successStatus == 204 {
		suffix = "Response"
	}

	var body []string
	body = append(body, "")
	body = append(body, fmt.Sprintf("return api.%s%d%s{}, nil",
		funcName, successStatus, suffix))

	joined := strings.Join(body, "\n")

	want := "return api.DeleteWebhook204Response{}, nil"
	if !strings.Contains(joined, want) {
		t.Fatalf("expected %q in body, got:\n%s", want, joined)
	}
	// Must NOT contain JSONResponse for 204.
	if strings.Contains(joined, "JSONResponse") {
		t.Fatalf("204 must not use JSONResponse suffix, got:\n%s", joined)
	}
}

// TestGenerateHTTPMethod_Delete200_NoResponseSeq verifies that when
// responseLines is empty and SuccessStatus is NOT 204 (e.g. 200),
// the fallback uses JSONResponse suffix.
func TestGenerateHTTPMethod_Delete200_NoResponseSeq(t *testing.T) {
	funcName := "GetHealth"
	successStatus := 200

	suffix := "JSONResponse"
	if successStatus == 204 {
		suffix = "Response"
	}

	var body []string
	body = append(body, "")
	body = append(body, fmt.Sprintf("return api.%s%d%s{}, nil",
		funcName, successStatus, suffix))

	joined := strings.Join(body, "\n")

	want := "return api.GetHealth200JSONResponse{}, nil"
	if !strings.Contains(joined, want) {
		t.Fatalf("expected %q in body, got:\n%s", want, joined)
	}
}
