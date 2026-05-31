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
//	<snakeID>.go                — <Symbol>Transitions var declaration
//	<snakeID>_can_transition.go — <Symbol>CanTransition guard function
//	<snakeID>_next_state.go     — <Symbol>NextState accessor function
//
// Splitting the prior single-file emit satisfies filefunc F1 on the
// statemachine package. id selects the filename stem so regeneration
// overwrites prior outputs in place. symbol (PascalCase of id, computed
// by the parser) is the only identifier used inside generated Go
// source; splicing id directly would emit unexported identifiers when
// the filename is lowercase — see BUG-002.
func writeStateFile(dir, id, symbol string, transMap map[string]map[string]string) error {
	base := strcase.ToSnake(id)

	files := map[string]string{
		base + ".go":                renderStateFile(id, symbol, transMap),
		base + "_can_transition.go": renderCanTransitionFile(id, symbol),
		base + "_next_state.go":     renderNextStateFile(id, symbol),
	}
	for name, src := range files {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
			return err
		}
	}
	return nil
}
