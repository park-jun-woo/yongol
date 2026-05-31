//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestSnakeCase_ZeroCov — Pascal/camel → snake (약어 런 처리)
package funcstub

import (
	"testing"
)

func TestSnakeCase_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"":           "",
		"Charge":     "charge",
		"ChargeCard": "charge_card",
		"getURL":     "get_url",
		"URLParser":  "url_parser",
		"parseID":    "parse_id",
	}
	for in, want := range cases {
		if got := snakeCase(in); got != want {
			t.Errorf("snakeCase(%q)=%q want %q", in, got, want)
		}
	}
}
