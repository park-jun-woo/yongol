//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestFrontendEnabled — nil manifest / enabled true/false/nil / 빈 블록 판정

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFrontendEnabled(t *testing.T) {
	boolPtr := func(b bool) *bool { return &b }

	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want bool
	}{
		{"nil manifest", &yongol.Fullstack{}, false},
		{
			"enabled nil + content",
			&yongol.Fullstack{Manifest: &manifest.ProjectConfig{Frontend: manifest.Frontend{Lang: "typescript"}}},
			true,
		},
		{
			"enabled true + content",
			&yongol.Fullstack{Manifest: &manifest.ProjectConfig{Frontend: manifest.Frontend{Enabled: boolPtr(true), Framework: "react"}}},
			true,
		},
		{
			"enabled false + content",
			&yongol.Fullstack{Manifest: &manifest.ProjectConfig{Frontend: manifest.Frontend{Enabled: boolPtr(false), Lang: "typescript"}}},
			false,
		},
		{
			"empty block",
			&yongol.Fullstack{Manifest: &manifest.ProjectConfig{Frontend: manifest.Frontend{}}},
			false,
		},
	}

	for _, c := range cases {
		if got := frontendEnabled(c.fs); got != c.want {
			t.Errorf("%s: frontendEnabled = %v, want %v", c.name, got, c.want)
		}
	}
}
