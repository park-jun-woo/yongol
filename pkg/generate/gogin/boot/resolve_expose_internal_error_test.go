//ff:func feature=gen-gogin type=test control=sequence topic=error-envelope
//ff:what resolveExposeInternalError — expose_internal_error 컴파일타임 기본값 (기본 false)

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveExposeInternalError(t *testing.T) {
	tval := true
	fval := false
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want bool
	}{
		{"nil fs", nil, false},
		{"nil manifest", &yongol.Fullstack{}, false},
		{"no error block", &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, false},
		{
			"error block nil flag",
			&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Error: &pmanifest.ErrorConfig{}}}},
			false,
		},
		{
			"explicit true",
			&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Error: &pmanifest.ErrorConfig{ExposeInternalError: &tval}}}},
			true,
		},
		{
			"explicit false",
			&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Error: &pmanifest.ErrorConfig{ExposeInternalError: &fval}}}},
			false,
		},
	}
	for _, c := range cases {
		if got := resolveExposeInternalError(c.fs); got != c.want {
			t.Errorf("%s: resolveExposeInternalError = %v, want %v", c.name, got, c.want)
		}
	}
}
