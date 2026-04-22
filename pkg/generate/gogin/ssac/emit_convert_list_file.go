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
) error {
	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature:   "service",
			Type:      "util",
			Control:   "iteration",
			Dimension: 1,
			Topic:     "response-serialize",
		},
		What: "convert" + name + "List — []db." + name + " → []api." + name + " 변환",
	}))
	sb.WriteString("package service\n\nimport (\n")
	sb.WriteString("\t\"" + modulePath + "/internal/api\"\n")
	sb.WriteString("\t\"" + modulePath + "/internal/db\"\n")
	sb.WriteString(")\n\n")
	sb.WriteString("func convert" + name + "List(rows []db." + name + ") []api." + name + " {\n")
	sb.WriteString("\tresult := make([]api." + name + ", len(rows))\n")
	sb.WriteString("\tfor i, row := range rows {\n")
	sb.WriteString("\t\tresult[i] = *convert" + name + "(row)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn result\n")
	sb.WriteString("}\n")

	fileName := fffile.EnsureUnique(fffile.FileNameForFunc("convert"+name+"List"), used)
	return fffile.WriteIfNotPreserved(filepath.Join(serviceDir, fileName), []byte(sb.String()))
}
