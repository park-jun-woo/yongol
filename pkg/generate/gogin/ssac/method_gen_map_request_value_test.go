//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what methodGen.mapRequestValue 단위 테스트 (path/query/body 분기 + format cast)

package ssac

import "testing"

func TestMethodGenMapRequestValue(t *testing.T) {
	g := &methodGen{
		PathParams:  map[string]bool{"id": true},
		QueryParams: map[string]queryParam{"limit": {GoType: "integer", IsRequired: true}},
		BodyFormats: map[string]string{"contact_email": "email"},
	}
	cases := []struct {
		name  string
		field string
		want  string
	}{
		{"path param", "id", "request.Id"},
		{"query param via accessor", "limit", "request.Params.Limit"},
		{"body field default", "title", "request.Body.Title"},
		{"body field with email cast", "contact_email", "string(request.Body.ContactEmail)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := g.mapRequestValue(tc.field); got != tc.want {
				t.Errorf("mapRequestValue(%q) = %q, want %q", tc.field, got, tc.want)
			}
		})
	}
}
