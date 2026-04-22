//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what walkPaths — paths 매핑을 순회하며 각 path 아래 operation 들을 색인
package openapi

import "gopkg.in/yaml.v3"

// walkPaths iterates the paths mapping and indexes operations under each path.
func walkPaths(paths *yaml.Node, idx *LineIndex) {
	if paths.Kind != yaml.MappingNode {
		return
	}
	for i := 0; i+1 < len(paths.Content); i += 2 {
		indexPathItem(paths.Content[i], paths.Content[i+1], idx)
	}
}
