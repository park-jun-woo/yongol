//ff:func feature=gen-gogin type=util control=sequence
//ff:what resolvePort — manifest.backend.http.port 값 결정 (미지정 시 8080)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

func resolvePort(fs *yongol.Fullstack) int {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.HTTP == nil {
		return defaultPort
	}
	p := fs.Manifest.Backend.HTTP.Port
	if p == 0 {
		return defaultPort
	}
	return p
}
