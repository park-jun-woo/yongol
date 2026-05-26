//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what S-71 — s71ExtractPrefix 단위 테스트 (변수 참조에서 prefix 추출)

package ssac

import "testing"

func TestS71ExtractPrefix(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"request.Name", "request"},
		{"course.ID", "course"},
		{"course", "course"},
		{"", ""},
		{`"literal"`, ""},
		{"123", ""},
		{"true", ""},
	}
	for _, tt := range tests {
		got := s71ExtractPrefix(tt.input)
		if got != tt.want {
			t.Errorf("s71ExtractPrefix(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
