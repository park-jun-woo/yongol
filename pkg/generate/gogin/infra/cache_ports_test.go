//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitCacheWrapper — cache adapter emit + 누락 포트/write 에러 검증
package infra

import (
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func cachePorts() []ssacmeta.Port {
	return []ssacmeta.Port{
		{Name: "CacheSet"},
		{Name: "CacheGet"},
		{Name: "CacheDelete"},
	}
}
