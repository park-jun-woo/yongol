//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what queueCfg — manifest.queue 섹션을 builtinBackend 로 래핑

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

func queueCfg(fs *yongol.Fullstack) builtinBackend {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Queue == nil {
		return builtinBackend{}
	}
	return builtinBackend{Present: true, Backend: fs.Manifest.Queue.Backend}
}
