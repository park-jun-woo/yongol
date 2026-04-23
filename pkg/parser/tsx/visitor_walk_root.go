//ff:func feature=tsx-parser type=parser control=iteration dimension=1
//ff:what (v *visitor).walkRoot — 루트 Module body 순회 진입점

package tsx

import "encoding/json"

// walkRoot decodes the root Module and visits its body.
func (v *visitor) walkRoot(root json.RawMessage) error {
	var m struct {
		Body []json.RawMessage `json:"body"`
	}
	if err := json.Unmarshal(root, &m); err != nil {
		return err
	}
	for _, node := range m.Body {
		v.walk(node)
	}
	return nil
}
