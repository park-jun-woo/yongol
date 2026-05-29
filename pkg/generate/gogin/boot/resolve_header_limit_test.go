//ff:func feature=gen-gogin type=test control=sequence
//ff:what resolveHeaderLimit — manifest.backend.http.header_limit 값 결정 (파싱 실패 시 1 MiB)

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveHeaderLimit(t *testing.T) {
	mk := func(raw string) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{HTTP: &pmanifest.HTTPConfig{HeaderLimit: raw}},
		}}
	}
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want int64
	}{
		{"nil fs", nil, defaultHeaderLimit},
		{"nil manifest", &yongol.Fullstack{}, defaultHeaderLimit},
		{"no http block", &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, defaultHeaderLimit},
		{"empty raw", mk(""), defaultHeaderLimit},
		{"unparseable raw", mk("not-a-size"), defaultHeaderLimit},
		{"valid 2MiB", mk("2MiB"), int64(2 << 20)},
	}
	for _, c := range cases {
		if got := resolveHeaderLimit(c.fs); got != c.want {
			t.Errorf("%s: resolveHeaderLimit = %d, want %d", c.name, got, c.want)
		}
	}
}
