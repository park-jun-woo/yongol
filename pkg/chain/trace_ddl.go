//ff:func feature=chain type=util control=iteration dimension=1
//ff:what traceDDL finds DDL tables referenced by SSaC sequences.
package chain

import (
	"log/slog"
	"strings"

	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func traceDDL(sf *ssac.ServiceFunc, tables []ddl.Table, specsDir string) []Link {
	tableSet := map[string]bool{}
	for _, t := range tables {
		tableSet[t.Name] = true
	}
	wanted := map[string]bool{}
	for _, seq := range sf.Sequences {
		if seq.Model == "" || seq.Type == "call" || seq.Type == "response" {
			continue
		}
		parts := strings.SplitN(seq.Model, ".", 2)
		if len(parts) < 2 {
			continue
		}
		name := inflection.Plural(toSnakeCase(parts[0]))
		if tableSet[name] {
			wanted[name] = true
		}
	}
	if len(wanted) == 0 {
		slog.Debug("chain.traceDDL: no DDL tables referenced by SSaC sequences", "operationId", sf.Name)
		return nil
	}
	var links []Link
	for _, name := range sortedStringKeys(wanted) {
		rel, line := findDDLTable(name, specsDir)
		links = append(links, Link{
			Kind:    "DDL",
			File:    rel,
			Line:    line,
			Summary: "CREATE TABLE " + name,
		})
	}
	return links
}
