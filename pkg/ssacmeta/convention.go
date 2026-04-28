//ff:type feature=ssacmeta type=model
//ff:what Convention — dynamic-port 네이밍 규약

package ssacmeta

// Convention is the dynamic-port naming & shape convention.
type Convention struct {
	Name        string  `yaml:"name"`
	Cardinality string  `yaml:"cardinality"`
	Params      []Field `yaml:"params,omitempty"`
	Returns     []Field `yaml:"returns,omitempty"`
}
