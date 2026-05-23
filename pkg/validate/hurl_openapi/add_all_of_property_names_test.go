//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what addAllOfPropertyNames — allOf 멤버의 property 이름 병합 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestAddAllOfPropertyNames(t *testing.T) {
	cases := []TestAddAllOfPropertyNamesCase{
		{
			name:     "nil_allof",
			allOf:    nil,
			existing: map[string]struct{}{},
			want:     map[string]struct{}{},
		},
		{
			name: "single_member_properties",
			allOf: openapi3.SchemaRefs{
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"id":   &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
						"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
					},
				}},
			},
			existing: map[string]struct{}{},
			want:     map[string]struct{}{"id": {}, "name": {}},
		},
		{
			name: "multiple_members_merged",
			allOf: openapi3.SchemaRefs{
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{"id": &openapi3.SchemaRef{Value: &openapi3.Schema{}}},
				}},
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{"email": &openapi3.SchemaRef{Value: &openapi3.Schema{}}},
				}},
			},
			existing: map[string]struct{}{},
			want:     map[string]struct{}{"id": {}, "email": {}},
		},
		{
			name: "nil_ref_skipped",
			allOf: openapi3.SchemaRefs{
				nil,
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{"ok": &openapi3.SchemaRef{Value: &openapi3.Schema{}}},
				}},
			},
			existing: map[string]struct{}{},
			want:     map[string]struct{}{"ok": {}},
		},
		{
			name: "nil_value_skipped",
			allOf: openapi3.SchemaRefs{
				{Value: nil},
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{"ok": &openapi3.SchemaRef{Value: &openapi3.Schema{}}},
				}},
			},
			existing: map[string]struct{}{},
			want:     map[string]struct{}{"ok": {}},
		},
		{
			name: "preserves_existing_entries",
			allOf: openapi3.SchemaRefs{
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{"new": &openapi3.SchemaRef{Value: &openapi3.Schema{}}},
				}},
			},
			existing: map[string]struct{}{"old": {}},
			want:     map[string]struct{}{"old": {}, "new": {}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runAddAllOfPropertyNames(t, c)
		})
	}
}
