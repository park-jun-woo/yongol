//ff:type feature=generate type=model
//ff:what State — generate 진입점에서 한 번 계산되는 불변 파생 상태

package prepared

// State is the immutable derivation computed at the top of Generate()
// and threaded through every emit helper. Emitters read already-
// resolved values from State instead of raw parsed manifest/SSaC/OpenAPI
// fields.
//
// Two invariants:
//
//  1. Optional subsystems (session/cache/file/queue) are represented as
//     nil in ActiveBackends when inactive. Emitters branch on nil
//     instead of "empty string vs default".
//
//  2. Resolved values are non-zero. For example Auth.Mode is always one
//     of "cookie"/"bearer"/"hybrid" — the Phase020 empty-string default
//     is applied inside prepared.AuthFor.
//
// State is infallible by contract: any configuration that cannot be
// resolved must have been rejected by validate as an ERROR before
// prepared.New() is called. The constructor must not panic and must
// not return an error.
type State struct {
	ActiveBackends ActiveBackends
	Auth           Auth
	Middlewares    []Middleware
	Routes         []Route

	// DomainAuth holds the per-domain resolved Auth for multi-domain
	// projects (Phase008). Each entry is the backend Auth with the domain's
	// auth_mode override applied (cookie/bearer/hybrid), so the route-group
	// wiring can pick the matching strict middleware per domain. nil for
	// single-site projects.
	DomainAuth map[string]Auth
}
