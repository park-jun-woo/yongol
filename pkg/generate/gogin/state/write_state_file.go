//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeStateFile — 상태 전이 3종(Transitions/CanTransition/NextState) 을 개별 파일로 기록

package state

import (
	"os"
	"path/filepath"

	"github.com/ettle/strcase"
)

// writeStateFile writes three sibling files for one statemachine:
//
//	<snakeID>.go                — <ID>Transitions var declaration
//	<snakeID>_can_transition.go — <ID>CanTransition guard function
//	<snakeID>_next_state.go     — <ID>NextState accessor function
//
// Splitting the prior single-file emit satisfies filefunc F1 on the
// statemachine package. The base name mirrors the diagram ID so
// regeneration overwrites prior outputs in place.
func writeStateFile(dir, id string, transMap map[string]map[string]string) error {
	base := strcase.ToSnake(id)

	files := map[string]string{
		base + ".go":                    renderStateFile(id, transMap),
		base + "_can_transition.go":     renderCanTransitionFile(id),
		base + "_next_state.go":         renderNextStateFile(id),
	}
	for name, src := range files {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
			return err
		}
	}
	return nil
}
