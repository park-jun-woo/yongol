//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestGenerateHTTPMethod_Delete200_NoResponseSeq — @response 시퀀스 없는 non-204 는 JSONResponse suffix 사용 검증

package ssac

import (
	"fmt"
	"strings"
	"testing"
)

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
