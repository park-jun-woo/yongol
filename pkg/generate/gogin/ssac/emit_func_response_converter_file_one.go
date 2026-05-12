//ff:func feature=gen-gogin type=generator control=sequence topic=response
//ff:what emitFuncResponseConverterFile — 단일 Func Response → api DTO converter 파일 emit

package ssac

import (
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// emitFuncResponseConverterFile writes a single
// internal/service/convert_<name>.go containing convert<Name> that maps a
// Func Response struct to the api DTO. Mirrors emitConvertFuncFile but
// imports the Func package instead of the db package.
func emitFuncResponseConverterFile(
	serviceDir, modulePath, name string,
	schema *openapi3.Schema,
	info funcRespInfo,
	spec *funcspec.FuncSpec,
	used map[string]bool,
) error {
	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature: "service",
			Type:    "util",
			Control: "sequence",
			Topic:   "response-serialize",
		},
		What: "convert" + name + " — " + info.PkgAlias + "." + name + " → *api." + name + " 변환",
	}))
	sb.WriteString("package service\n\nimport (\n")
	sb.WriteString("\t\"" + modulePath + "/internal/api\"\n")
	if info.ImportPath != "" {
		sb.WriteString("\t\"" + info.ImportPath + "\"\n")
	}
	sb.WriteString(")\n\n")
	writeFuncResponseConvertFunc(&sb, name, schema, info, spec)

	fileName := fffile.EnsureUnique(fffile.FileNameForFunc("convert"+name), used)
	return fffile.WriteIfNotPreserved(filepath.Join(serviceDir, fileName), []byte(sb.String()))
}
