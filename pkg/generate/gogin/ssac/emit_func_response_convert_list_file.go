//ff:func feature=gen-gogin type=generator control=sequence topic=response
//ff:what emitFuncResponseConvertListFile — Func Response convert<Name>List 파일을 1파일 1func 로 emit

package ssac

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

// emitFuncResponseConvertListFile writes
// internal/service/convert_<name>_list.go with a single
// convert<Name>List helper that maps a slice of Func package types to a
// slice of api DTOs by delegating to convert<Name>. Mirrors
// emitConvertListFile but imports the Func package instead of the db
// package.
func emitFuncResponseConvertListFile(
	serviceDir, modulePath, name string,
	info funcRespInfo,
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
		What: "convert" + base + "List — []" + info.PkgAlias + "." + name + " → []api." + name + " 변환",
	}))
	sb.WriteString("package service\n\nimport (\n")
	sb.WriteString("\t" + apiImportLine(modulePath, dg.ApiSuffix) + "\n")
	if info.ImportPath != "" {
		sb.WriteString("\t\"" + info.ImportPath + "\"\n")
	}
	sb.WriteString(")\n\n")
	sb.WriteString("func convert" + base + "List(rows []" + info.PkgAlias + "." + name + ") ([]api." + name + ", error) {\n")
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
