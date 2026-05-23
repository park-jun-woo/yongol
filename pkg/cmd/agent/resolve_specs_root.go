//ff:func feature=agent type=helper control=selection
//ff:what resolveSpecsRoot — 파일 절대 경로와 레이어로부터 specs 루트 디렉토리 추론

package agent

import "path/filepath"

// resolveSpecsRoot derives the specs root directory from an absolute file path and its layer.
func resolveSpecsRoot(absPath string, l layer) string {
	switch l {
	case layerOpenAPI:
		return filepath.Dir(filepath.Dir(absPath))
	case layerRego:
		return filepath.Dir(filepath.Dir(absPath))
	case layerHurl:
		return filepath.Dir(filepath.Dir(absPath))
	default:
		return filepath.Dir(absPath)
	}
}
