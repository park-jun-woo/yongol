//ff:type feature=ssacmeta type=model
//ff:what Port — 단일 DB-access port 선언

package ssacmeta

// Port is a single DB-access port declared by an ssac package.
type Port struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	When        string    `yaml:"when"`
	UsedBy      []string  `yaml:"used_by,omitempty"`
	Query       QuerySpec `yaml:"query"`
}
