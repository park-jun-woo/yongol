//ff:func feature=rule type=loader control=sequence topic=catalog
//ff:what MustLoad — Load 래퍼 (실패 시 log.Fatal), CLI 엔트리에서 사용
package catalog

import "log"

// MustLoad returns the embedded catalog or calls log.Fatal on parse error.
// Intended for CLI entry points where recovery is not meaningful.
func MustLoad() *Catalog {
	c, err := Load()
	if err != nil {
		log.Fatalf("yongol: %v", err)
	}
	return c
}
