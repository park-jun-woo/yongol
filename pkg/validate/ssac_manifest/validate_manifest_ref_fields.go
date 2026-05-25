//ff:func feature=validate type=helper control=iteration dimension=1 topic=ssac-manifest
//ff:what validateManifestRefFields — 단일 시퀀스의 manifest.* 필드 참조를 검증

package ssac_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// validateManifestRefFields checks each field in a @response sequence for
// manifest.* references and validates them against the manifest config.
func validateManifestRefFields(funcName, fileName string, seq ssacparser.Sequence, fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if seq.Type != "response" || len(seq.Fields) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for fieldName, fieldVal := range seq.Fields {
		if !strings.HasPrefix(fieldVal, "manifest.") {
			continue
		}
		diags = append(diags, validateOneManifestRef(funcName, fileName, seq.Line, fieldName, fieldVal, fs.Manifest)...)
	}
	return diags
}
