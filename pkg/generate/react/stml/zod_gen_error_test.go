//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what zodGenError.Error — 빈 OperationID/Field 의 <unknown> 폴백 분기 검증

package stml

import (
	"strings"
	"testing"
)

func TestZodGenError_Error(t *testing.T) {
	cases := []struct {
		name    string
		err     zodGenError
		wantSub []string
	}{
		{
			name:    "fully populated",
			err:     zodGenError{OperationID: "CreateThing", Field: "title", Type: "weirdtype"},
			wantSub: []string{"CreateThing", "title", "weirdtype"},
		},
		{
			name:    "empty op falls back to unknown",
			err:     zodGenError{OperationID: "", Field: "title", Type: "weirdtype"},
			wantSub: []string{"<unknown>", "title", "weirdtype"},
		},
		{
			name:    "empty field falls back to unknown",
			err:     zodGenError{OperationID: "CreateThing", Field: "", Type: "weirdtype"},
			wantSub: []string{"CreateThing", "<unknown>", "weirdtype"},
		},
		{
			name:    "both empty fall back to unknown",
			err:     zodGenError{},
			wantSub: []string{"<unknown>"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			msg := tc.err.Error()
			for _, sub := range tc.wantSub {
				if !strings.Contains(msg, sub) {
					t.Errorf("Error() = %q, missing %q", msg, sub)
				}
			}
		})
	}
}
