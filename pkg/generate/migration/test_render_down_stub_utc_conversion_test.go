//ff:func feature=migration type=test control=sequence
//ff:what TestRenderDownStub_UTCConversion — non-UTC 입력도 header 에 UTC 로 직렬화되는지 확인
package migration

import (
	"strings"
	"testing"
	"time"
)

// TestRenderDownStub_UTCConversion verifies that non-UTC input times are
// normalised to UTC before being embedded into the stub header.
func TestRenderDownStub_UTCConversion(t *testing.T) {
	// Non-UTC input must be normalised to UTC in the header.
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		t.Skipf("Asia/Seoul tzdata unavailable: %v", err)
	}
	kst := time.Date(2026, 4, 24, 21, 34, 56, 0, loc) // == 12:34:56Z
	body := RenderDownStub("v1", kst)
	if !strings.Contains(body, "2026-04-24T12:34:56Z") {
		t.Errorf("expected UTC timestamp in body, got:\n%s", body)
	}
}
