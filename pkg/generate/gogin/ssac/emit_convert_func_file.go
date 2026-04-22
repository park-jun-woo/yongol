//ff:func feature=gen-gogin type=generator control=sequence
//ff:what emitConvertFuncFile — 단일 스키마 convert<Name> 함수를 1파일 1func 로 emit

package ssac

import (
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

// emitConvertFuncFile writes internal/service/convert_<name>.go containing a
// single convert<Name> function (db row → api DTO). The file carries a
// complete //ff:func + //ff:what header so filefunc A1/A3 pass without any
// follow-up annotator pass.
//
// The caller (emitAllConverterFiles) is responsible for ensuring used
// is non-nil; EnsureUnique is invoked so repeated schemas (theoretically
// impossible, but defensive) never collide.
func emitConvertFuncFile(
	serviceDir, modulePath, name string,
	schema *openapi3.Schema,
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
		What: "convert" + name + " — db." + name + " row → *api." + name + " 변환",
	}))
	sb.WriteString("package service\n\nimport (\n")
	sb.WriteString("\t\"" + modulePath + "/internal/api\"\n")
	sb.WriteString("\t\"" + modulePath + "/internal/db\"\n")
	sb.WriteString(")\n\n")
	writeConvertFunc(&sb, name, schema)

	fileName := fffile.EnsureUnique(fffile.FileNameForFunc("convert"+name), used)
	return fffile.WriteIfNotPreserved(filepath.Join(serviceDir, fileName), []byte(sb.String()))
}
