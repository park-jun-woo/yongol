//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-manifest
//ff:what XNS-80 — SSaC @response manifest.* 참조가 manifest.yaml 에 실제 존재하는지 교차 검증

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xns80ManifestRef validates XNS-80: every manifest.* reference in SSaC
// @response fields must resolve to an actual value in manifest.yaml.
func xns80ManifestRef(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, sf := range fs.ServiceFuncs {
		for _, seq := range sf.Sequences {
			diags = append(diags, validateManifestRefFields(sf.Name, sf.FileName, seq, fs)...)
		}
	}
	return diags
}
