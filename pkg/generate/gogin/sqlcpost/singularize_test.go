//ff:func feature=gen-gogin type=test control=branch topic=sqlc-post
//ff:what TestSingularize — ies/es-sibilant/s/무변환 각 분기 검증

package sqlcpost

import "testing"

func TestSingularize(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"categories", "category"}, // ies -> y
		{"buses", "bus"},           // ses -> drop es
		{"boxes", "box"},           // xes
		{"quizes", "quiz"},         // zes
		{"matches", "match"},       // ches
		{"dishes", "dish"},         // shes
		{"users", "user"},          // s -> drop
		{"data", "data"},           // no change
		{"", ""},                   // empty
	}
	for _, c := range cases {
		if got := singularize(c.in); got != c.want {
			t.Errorf("singularize(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
