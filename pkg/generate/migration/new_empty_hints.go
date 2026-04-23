//ff:func feature=migration type=util control=sequence topic=migration-hints
//ff:what newEmptyHints — 맵 필드를 초기화한 빈 Hints 구조체 반환
package migration

// newEmptyHints returns a Hints with its map fields initialised.
func newEmptyHints() *Hints {
	return &Hints{
		Casts:            map[colKey]string{},
		Backfills:        map[colKey]string{},
		DataMigrations:   map[string]string{},
		AllowDestructive: map[string]bool{},
	}
}
