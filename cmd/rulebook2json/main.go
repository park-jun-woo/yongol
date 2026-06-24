//ff:func feature=rule type=command control=sequence topic=catalog
//ff:what rulebook.md(내장 카탈로그)를 reins 퀘스트가 시드로 쓰는 rules.json 으로 변환한다. 손파싱 드리프트를 피하려 catalog.Load 를 그대로 재사용 — 활성 규칙(폐기 제외)만, rulebook 순서 보존. 인자로 출력 경로(없으면 stdout).

// Command rulebook2json converts the embedded rulebook.md catalog into the
// rules.json the reins rule-migration quest seeds from. It reuses the existing
// pkg/rule/catalog loader verbatim (no hand re-parsing → no drift), emitting only
// the active rules (the Deprecated section is already skipped by Load) in rulebook
// order. With one positional argument it writes there; otherwise stdout.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

func main() {
	cat, err := catalog.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "load rulebook catalog:", err)
		os.Exit(1)
	}

	rules := toOutRules(cat)

	data, err := json.MarshalIndent(rules, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "marshal rules:", err)
		os.Exit(1)
	}
	data = append(data, '\n')

	if len(os.Args) > 1 {
		if err := os.WriteFile(os.Args[1], data, 0o644); err != nil {
			fmt.Fprintln(os.Stderr, "write:", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "wrote %d active rules → %s\n", len(rules), os.Args[1])
		return
	}
	os.Stdout.Write(data)
}
