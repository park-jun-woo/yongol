//ff:func feature=gen-gogin type=generator control=sequence
//ff:what emitConvertListFile — 단일 스키마 convert<Name>List 함수를 1파일 1func 로 emit

package ssac

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

// emitConvertListFile writes internal/service/convert_<name>_list.go with a
// single convert<Name>List helper that maps a slice of db rows to a slice of
// api DTOs by delegating to convert<Name>. Emitted alongside
// emitConvertFuncFile when the target service depends on the schema.
func emitConvertListFile(
	serviceDir, modulePath, name string,
	used map[string]bool,
	dg domainGen,
) error {
	base := dg.FuncPrefix + name
	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature:   "service",
			Type:      "util",
			Control:   "iteration",
			Dimension: 1,
			Topic:     "response-serialize",
		},
		What: "convert" + base + "List — []db." + name + " → []api." + name + " 변환",
	}))
	sb.WriteString("package service\n\nimport (\n")
	sb.WriteString("\t" + apiImportLine(modulePath, dg.ApiSuffix) + "\n")
	sb.WriteString("\t\"" + modulePath + "/internal/db\"\n")
	sb.WriteString(")\n\n")
	// convert<Model> now returns (*api.X, error) because JSONB unmarshal
	// can fail (BUG-005). The list variant must propagate that error
	// instead of swallowing it — unmarshal errors are 500-class bugs
	// and tx rollback only works when the handler sees them.
	sb.WriteString("func convert" + base + "List(rows []db." + name + ") ([]api." + name + ", error) {\n")
	sb.WriteString("\tresult := make([]api." + name + ", len(rows))\n")
	sb.WriteString("\tfor i, row := range rows {\n")
	sb.WriteString("\t\titem, err := convert" + base + "(row)\n")
	sb.WriteString("\t\tif err != nil {\n")
	sb.WriteString("\t\t\treturn nil, err\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t\tresult[i] = *item\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn result, nil\n")
	sb.WriteString("}\n")

	fileName := fffile.EnsureUnique(fffile.FileNameForFunc("convert"+base+"List"), used)
	return fffile.WriteIfNotPreserved(filepath.Join(serviceDir, fileName), []byte(sb.String()))
}
