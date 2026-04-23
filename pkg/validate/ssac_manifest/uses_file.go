//ff:func feature=validate type=util control=iteration dimension=2 topic=config-check
//ff:what usesFile — SSaC 에 @call file.* 또는 @call storage.* 호출이 있는지 확인

package ssac_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// usesFile reports whether any SSaC service func issues an @call whose
// model starts with "file." or "storage.". Both prefixes map to the same
// manifest file backend (local/s3) in yongol codegen, so either one
// signals that file.backend configuration is required.
func usesFile(fs *yongol.Fullstack) bool {
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" {
				continue
			}
			if strings.HasPrefix(seq.Model, "file.") || strings.HasPrefix(seq.Model, "storage.") {
				return true
			}
		}
	}
	return false
}
