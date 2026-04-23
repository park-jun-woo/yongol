//ff:func feature=generate type=util control=sequence
//ff:what routesFor — OpenAPI 기반 라우트 해석 (placeholder)

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// routesFor is a Phase001 placeholder returning nil. Routes stay in
// the existing emitters for now; Stage 5 migrates them onto
// prepared.Route once the emit signatures switch to prepared.State.
func routesFor(fs *yongol.Fullstack) []Route {
	_ = fs
	return nil
}
