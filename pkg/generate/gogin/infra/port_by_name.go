//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what portByName — active port 목록에서 name 과 일치하는 port 를 찾음

package infra

import "github.com/park-jun-woo/yongol/pkg/ssacmeta"

// portByName locates a port from the active subset by its interface.yaml
// name (e.g. "CacheSet"). Returns nil when the port is inactive under the
// current manifest — callers must handle the missing port (typically by
// skipping the method that depends on it).
func portByName(active []ssacmeta.Port, name string) *ssacmeta.Port {
	for i := range active {
		if active[i].Name == name {
			return &active[i]
		}
	}
	return nil
}
