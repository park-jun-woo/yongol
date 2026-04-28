//ff:type feature=ssacmeta type=model
//ff:what PackageInterface — ssac 패키지의 interface.yaml AST 루트

package ssacmeta

// PackageInterface is the parsed representation of a single ssac package's
// interface.yaml. One instance per DB-using ssac package (cache / session /
// queue / auth / authz …).
type PackageInterface struct {
	Version          int    `yaml:"version"`
	Package          string `yaml:"package"`
	Description      string `yaml:"description"`
	Ports            []Port `yaml:"ports"`
	PortsMeta        *Meta  `yaml:"ports_meta,omitempty"`
	CanonicalDDL     string `yaml:"canonical_ddl,omitempty"`
	CanonicalQueries string `yaml:"canonical_queries,omitempty"`

	// SourcePath is populated by the loader so diagnostics can reference
	// the originating file.
	SourcePath string `yaml:"-"`
}
