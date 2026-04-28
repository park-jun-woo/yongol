//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what sessionCfg — manifest.session 섹션을 builtinBackend 로 래핑

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

func sessionCfg(fs *yongol.Fullstack) builtinBackend {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Session == nil {
		return builtinBackend{}
	}
	return builtinBackend{Present: true, Backend: fs.Manifest.Session.Backend}
}
