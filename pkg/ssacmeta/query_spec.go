//ff:type feature=ssacmeta type=model
//ff:what QuerySpec — Port 를 뒷받침하는 sqlc 쿼리 스펙

package ssacmeta

// QuerySpec describes the sqlc query that backs a Port.
type QuerySpec struct {
	Cardinality string  `yaml:"cardinality"` // one | many | exec
	Params      []Field `yaml:"params,omitempty"`
	Returns     []Field `yaml:"returns,omitempty"`
}
