//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what collectFromValueMap — request. 접두어 수집/비request 무시 검증

package openapi_ssac

import "testing"

func TestCollectFromValueMap(t *testing.T) {
	t.Run("empty map does nothing", func(t *testing.T) {
		fields := map[string]bool{}
		collectFromValueMap(fields, nil)
		if len(fields) != 0 {
			t.Errorf("expected empty, got %v", fields)
		}
	})

	t.Run("collects request-prefixed values", func(t *testing.T) {
		fields := map[string]bool{}
		m := map[string]string{
			"status":  "request.Status",
			"name":    "course.Name",
			"user_id": "request.UserID",
		}
		collectFromValueMap(fields, m)
		if len(fields) != 2 {
			t.Fatalf("expected 2, got %d: %v", len(fields), fields)
		}
		if !fields["Status"] || !fields["UserID"] {
			t.Errorf("missing expected: %v", fields)
		}
	})
}
