//ff:func feature=projectconfig type=test control=iteration dimension=1
//ff:what DomainConfig.ResolvedAllowOrigins — domain override / backend 상속 / nil 분기 검증

package manifest

import (
	"reflect"
	"testing"
)

func TestDomainConfigResolvedAllowOrigins(t *testing.T) {
	backend := &CORSConfig{AllowOrigins: []string{"https://app.example.com"}}
	domainCORS := &CORSConfig{AllowOrigins: []string{"https://admin.example.com"}}

	cases := []struct {
		name    string
		cors    *CORSConfig
		backend *CORSConfig
		want    []string
	}{
		{"domain override", domainCORS, backend, []string{"https://admin.example.com"}},
		{"inherit backend", nil, backend, []string{"https://app.example.com"}},
		{"both nil", nil, nil, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			d := DomainConfig{CORS: c.cors}
			if got := d.ResolvedAllowOrigins(c.backend); !reflect.DeepEqual(got, c.want) {
				t.Errorf("ResolvedAllowOrigins = %v, want %v", got, c.want)
			}
		})
	}
}
