//ff:func feature=validate type=helper control=sequence topic=ssac-manifest
//ff:what validateOneManifestRef — 단일 manifest.* 참조의 존재·유효성 검증

package ssac_manifest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// validateOneManifestRef validates a single manifest.* field reference.
func validateOneManifestRef(funcName, fileName string, line int, fieldName, fieldVal string, mf *manifest.ProjectConfig) []diagnostic.Diagnostic {
	refPath := strings.TrimPrefix(fieldVal, "manifest.")
	if !isKnownRefPath(refPath) {
		return []diagnostic.Diagnostic{{
			File:        fileName,
			Line:        line,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[XNS-80] @response field %q references unknown manifest path %q", fieldName, fieldVal),
			Advice:      fmt.Sprintf("Supported manifest.* paths: %s", strings.Join(manifest.KnownRefPaths(), ", ")),
			OperationID: funcName,
		}}
	}
	if _, ok := manifest.ResolveRef(mf, refPath); !ok {
		return []diagnostic.Diagnostic{{
			File:        fileName,
			Line:        line,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[XNS-80] @response field %q references %q but the value is missing or invalid in manifest.yaml", fieldName, fieldVal),
			Advice:      fmt.Sprintf("Ensure manifest.yaml declares a valid value at the path: %s", refPath),
			OperationID: funcName,
		}}
	}
	return nil
}
