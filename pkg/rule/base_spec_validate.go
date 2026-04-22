//ff:func feature=rule type=accessor
//ff:what Validate — BaseSpec 필드 건전성 검증 (Rule/Level non-empty)
package rule

import "errors"

// Validate checks that BaseSpec carries a non-empty Rule ID and Level.
// Message may be empty (some rules format it at runtime).
func (s BaseSpec) Validate() error {
	if s.Rule == "" {
		return errors.New("rule: BaseSpec.Rule is empty")
	}
	if s.Level == "" {
		return errors.New("rule: BaseSpec.Level is empty")
	}
	return nil
}
