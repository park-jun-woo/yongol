//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithFuncs(funcs ...ssac.ServiceFunc) *yongol.Fullstack {
	return &yongol.Fullstack{ServiceFuncs: funcs}
}
