//ff:func feature=generate type=util control=sequence
//ff:what fileBackendFor — manifest + SSaC 사용 여부로 file 활성 판정 및 기본값 해석

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// fileBackendFor returns non-nil iff the file subsystem is in use.
// Default "local" applied when SSaC uses file but manifest is silent.
func fileBackendFor(fs *yongol.Fullstack) *File {
	if manifestDeclaresFile(fs) {
		return &File{Backend: fs.Manifest.File.Backend}
	}
	if ssacUsesFileCalls(fs) {
		return &File{Backend: "local"}
	}
	return nil
}
