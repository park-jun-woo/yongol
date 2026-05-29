//ff:func feature=validate type=test-helper control=sequence topic=domain-security
//ff:what cleanupFS — 테스트용 임시 디렉토리 삭제
package domain_security

import (
	"os"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// cleanupFS removes the temporary directory used for testing.
func cleanupFS(fs *yongol.Fullstack) {
	os.RemoveAll(fs.SpecsDir)
}
