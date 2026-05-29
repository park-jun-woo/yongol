//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isNonPointerInput 단위 테스트 (리터럴 / path / required query / required body 분기)

package ssac

import "testing"

func TestIsNonPointerInput(t *testing.T) {
	g := &methodGen{
		PathParams: map[string]bool{"id": true},
		QueryParams: map[string]queryParam{
			"limit":  {IsRequired: true},
			"cursor": {IsRequired: false},
		},
		BodyRequiredFields: map[string]bool{"name": true},
	}
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"literal", `"x"`, true},
		{"number literal", "5", true},
		{"non-request var", "user.ID", false},
		{"required path param", "request.id", true},
		{"required query", "request.limit", true},
		{"optional query", "request.cursor", false},
		{"required body field", "request.name", true},
		{"unknown body field optional", "request.bio", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := g.isNonPointerInput(tc.in); got != tc.want {
				t.Errorf("isNonPointerInput(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
