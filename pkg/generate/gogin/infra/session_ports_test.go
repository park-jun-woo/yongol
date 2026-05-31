//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitSessionWrapper — session adapter emit + 누락 포트/write 에러 검증
package infra

import (
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func sessionPorts() []ssacmeta.Port {
	return []ssacmeta.Port{
		{Name: "SessionSet"},
		{Name: "SessionGet"},
		{Name: "SessionDelete"},
	}
}
