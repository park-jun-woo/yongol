//ff:type feature=ssacmeta type=model
//ff:what Meta — dynamic-port 패키지(authz 등)의 ports_meta 정의

package ssacmeta

// Meta holds ports_meta (for dynamic-port packages like authz) — the
// concrete port list is not known ahead of time; instead the meta rule
// describes the naming convention and shape.
type Meta struct {
	Rule        string      `yaml:"rule"`
	Description string      `yaml:"description"`
	Convention  *Convention `yaml:"convention,omitempty"`
}
