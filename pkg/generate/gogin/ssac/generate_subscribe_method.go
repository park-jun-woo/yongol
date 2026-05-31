//ff:func feature=gen-gogin type=generator control=iteration dimension=3
//ff:what generateSubscribeMethod — @subscribe → Server method

package ssac

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ettle/strcase"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateSubscribeMethod writes a subscribe handler method.
// Signature: func (server *Server) OnXxx(ctx context.Context, msg []byte) error
func generateSubscribeMethod(sf ssacparser.ServiceFunc, fs *yongol.Fullstack, serviceDir, modulePath string) error {
	useTx := needsTransaction(sf)
	g := newMethodGen(fs.OpenAPIDoc, sf, modulePath, useTx, fs.ProjectFuncSpecs, fs.YongolPkgSpecs, tracingWrapCalls(fs), fs.StateDiagrams, collectOwnerships(fs), fs.DDLTables, fs.SQLcQueries, fs.Manifest)
	g.IsSubscribe = true

	var imports []string
	var body []string

	imports = append(imports, `"context"`, `"encoding/json"`, `"log/slog"`)

	// Handler entry DEBUG log (Phase012 AutoLogInsert 1단계) — topic tag
	topic := ""
	if sf.Subscribe != nil {
		topic = sf.Subscribe.Topic
	}
	body = append(body, fmt.Sprintf(`slog.DebugContext(ctx, "handler entry", "op", %q, "topic", %q)`, sf.Name, topic))

	// message unmarshal
	if sf.Subscribe != nil {
		body = append(body, "var message struct {")
		for _, st := range sf.Structs {
			if st.Name == sf.Subscribe.MessageType {
				for _, f := range st.Fields {
					body = append(body, fmt.Sprintf("\t%s %s `json:\"%s\"`", f.Name, f.Type, f.Name))
				}
			}
		}
		body = append(body, "}")
		body = append(body, "if err := json.Unmarshal(msg, &message); err != nil { return err }")
	}

	if useTx {
		imports = append(imports, `"fmt"`)
		imports, body = appendSubscribeTxBeginLines(imports, body, sf.Name)
	}

	var postCommit []string
	for i, seq := range sf.Sequences {
		if seq.Type == "response" {
			continue
		}
		var next *ssacparser.Sequence
		if i+1 < len(sf.Sequences) {
			next = &sf.Sequences[i+1]
		}
		lines, imp, isPost, err := g.buildSequence(seq, next)
		if err != nil {
			return fmt.Errorf("generateSubscribeMethod %s: %w", sf.Name, err)
		}
		imports = append(imports, imp...)
		// db import needed when sqlcArgs generates db.XXXParams (2+ inputs)
		if isCRUD(seq.Type) && len(seq.Inputs) > 1 {
			imports = append(imports, fmt.Sprintf(`"%s/internal/db"`, modulePath))
		}
		if isPost {
			postCommit = append(postCommit, lines...)
		} else {
			body = append(body, "")
			body = append(body, lines...)
		}
	}

	if useTx {
		body = append(body, "", "if err := tx.Commit(ctx); err != nil { return fmt.Errorf(\"commit: %w\", err) }")
	}
	if len(postCommit) > 0 {
		body = append(body, "")
		body = append(body, postCommit...)
	}
	body = append(body, "", "return nil")

	sig := fmt.Sprintf("func (server *Server) %s(ctx context.Context, msg []byte) error", sf.Name)

	// dedup imports
	seen := make(map[string]bool)
	var deduped []string
	for _, imp := range imports {
		if imp != "" && !seen[imp] {
			seen[imp] = true
			deduped = append(deduped, imp)
		}
	}
	sort.Strings(deduped)

	control := ffannot.DetectControl(body)
	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature:   "service",
			Type:      "handler",
			Control:   control,
			Dimension: 1,
			Topic:     "subscribe",
		},
		What: sf.Name + " — queue subscribe handler (topic " + topic + ")",
	})

	var sb strings.Builder
	sb.WriteString(header)
	sb.WriteString("package service\n\nimport (\n")
	for _, imp := range deduped {
		sb.WriteString("\t" + imp + "\n")
	}
	sb.WriteString(")\n\n")
	sb.WriteString(sig + " {\n")
	for _, line := range body {
		if line == "" {
			sb.WriteString("\n")
		} else {
			sb.WriteString("\t" + line + "\n")
		}
	}
	sb.WriteString("}\n")

	fileName := strcase.ToSnake(sf.Name) + ".go"
	return fffile.WriteIfNotPreserved(filepath.Join(serviceDir, fileName), []byte(sb.String()))
}
