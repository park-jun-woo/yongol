//ff:type feature=ssacmeta type=model
//ff:what PackageInterface / Port / Query — ssac 패키지의 interface.yaml AST

package ssacmeta

// PackageInterface is the parsed representation of a single ssac package's
// interface.yaml. One instance per DB-using ssac package (cache / session /
// queue / auth / authz …).
type PackageInterface struct {
	Version          int     `yaml:"version"`
	Package          string  `yaml:"package"`
	Description      string  `yaml:"description"`
	Ports            []Port  `yaml:"ports"`
	PortsMeta        *Meta   `yaml:"ports_meta,omitempty"`
	CanonicalDDL     string  `yaml:"canonical_ddl,omitempty"`
	CanonicalQueries string  `yaml:"canonical_queries,omitempty"`

	// SourcePath is populated by the loader so diagnostics can reference
	// the originating file.
	SourcePath string `yaml:"-"`
}

// Port is a single DB-access port declared by an ssac package.
type Port struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	When        string    `yaml:"when"`
	UsedBy      []string  `yaml:"used_by,omitempty"`
	Query       QuerySpec `yaml:"query"`
}

// QuerySpec describes the sqlc query that backs a Port.
type QuerySpec struct {
	Cardinality string   `yaml:"cardinality"` // one | many | exec
	Params      []Field  `yaml:"params,omitempty"`
	Returns     []Field  `yaml:"returns,omitempty"`
}

// Field is a single named/typed field in params or returns.
type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// Meta holds ports_meta (for dynamic-port packages like authz) — the
// concrete port list is not known ahead of time; instead the meta rule
// describes the naming convention and shape.
type Meta struct {
	Rule        string      `yaml:"rule"`
	Description string      `yaml:"description"`
	Convention  *Convention `yaml:"convention,omitempty"`
}

// Convention is the dynamic-port naming & shape convention.
type Convention struct {
	Name        string  `yaml:"name"`
	Cardinality string  `yaml:"cardinality"`
	Params      []Field `yaml:"params,omitempty"`
	Returns     []Field `yaml:"returns,omitempty"`
}
