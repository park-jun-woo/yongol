//ff:type feature=migration type=model
//ff:what Options — Generate 튜닝 옵션 (버전/시간)
package migration

import "time"

// Options tunes Generate.
type Options struct {
	YongolVersion string
	Now           time.Time // for deterministic tests; zero => time.Now()
}
