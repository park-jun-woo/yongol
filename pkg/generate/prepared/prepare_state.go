//ff:func feature=generate type=util control=sequence
//ff:what New — Fullstack → prepared.State 변환 진입점 (infallible)

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// New is the first step of Generate(). It computes every derived value
// emitters need — active backends, resolved auth, middleware list,
// route table — from the parsed Fullstack.
//
// New is infallible by contract: any configuration that cannot be
// resolved must have been rejected by validate as an ERROR before
// Generate() is called. The function must not panic and must not
// return an error.
func New(fs *yongol.Fullstack) State {
	return State{
		ActiveBackends: activeBackendsFor(fs),
		Auth:           AuthFor(fs),
		Middlewares:    middlewaresFor(fs),
		Routes:         routesFor(fs),
	}
}
