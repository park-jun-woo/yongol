//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what mergeMethodBlock — 새 method 블록을 기존 pathBlocks에 병합

package agent

func mergeMethodBlock(pathBlocks map[string]any, pathKey string, block map[string]any) {
	existing, ok := pathBlocks[pathKey]
	if !ok {
		pathBlocks[pathKey] = block
		return
	}
	existingMap, ok := existing.(map[string]any)
	if !ok {
		pathBlocks[pathKey] = block
		return
	}
	for method, detail := range block {
		existingMap[method] = detail
	}
	pathBlocks[pathKey] = existingMap
}
