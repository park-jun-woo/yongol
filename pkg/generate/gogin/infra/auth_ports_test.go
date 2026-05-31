//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitAuthWrapper — auth RefreshStore adapter 6파일 emit + 누락/mkdir/write 에러 검증
package infra

import (
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func authPorts() []ssacmeta.Port {
	return []ssacmeta.Port{
		{Name: "RefreshTokenInsert"},
		{Name: "RefreshTokenConsume"},
		{Name: "RefreshTokenCheckReuse"},
		{Name: "RefreshTokenRevoke"},
		{Name: "RefreshTokenRevokeAll"},
	}
}
