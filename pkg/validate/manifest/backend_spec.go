//ff:type feature=validate type=model topic=manifest-infra
//ff:what backendSpec — XNC/XNS/XNQ-90 검증에 필요한 항목 명세

package manifest

// backendSpec declares everything validateBuiltinBackend needs to check
// one backend. Defined as a struct rather than positional args so future
// additions (e.g. an optional second DDL table) don't break call sites.
type backendSpec struct {
	Pkg            string
	Cfg            builtinBackend
	RequireDDL     string
	RequireQueries []string
	RuleID         string
}
