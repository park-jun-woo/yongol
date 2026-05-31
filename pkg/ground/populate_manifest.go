//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateManifest — manifest.yaml에서 middleware, claims, roles, queue 설정 추출
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func populateManifest(g *rule.Ground, fs *yongol.Fullstack) {
	if fs.Manifest == nil {
		return
	}
	middleware := make(rule.StringSet)
	for _, m := range fs.Manifest.Backend.Middleware {
		middleware[m] = true
	}
	g.Lookup["Manifest.middleware"] = middleware

	if fs.Manifest.Backend.Auth != nil {
		populateManifestAuth(g, fs.Manifest.Backend.Auth)
	}

	if fs.Manifest.Queue != nil && fs.Manifest.Queue.Backend != "" {
		g.Config["queue.backend"] = true
	}
	if len(fs.Manifest.Backend.Middleware) > 0 {
		g.Config["backend.middleware"] = true
	}
}
