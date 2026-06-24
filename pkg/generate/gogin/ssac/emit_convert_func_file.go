//ff:func feature=gen-gogin type=generator control=sequence
//ff:what emitConvertFuncFile — 단일 스키마 convert<Name> 함수를 1파일 1func 로 emit

package ssac

import (
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
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
	ddlTables []ddl.Table,
	used map[string]bool,
	dg domainGen,
) error {
	funcName := "convert" + dg.FuncPrefix + name
	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature: "service",
			Type:    "util",
			Control: "sequence",
			Topic:   "response-serialize",
		},
		What: funcName + " — db." + name + " row → *api." + name + " 변환",
	}))
	// encoding/json is imported unconditionally — even for schemas that
	// have no JSONB fields the convert function signature now returns an
	// error so callers use one pattern. The Go compiler forbids unused
	// imports, but encoding/json is always used when at least one JSONB
	// field exists; for JSONB-less schemas the import is still referenced
	// by the closing comment block? No — it would be unused. We only add
	// the import when at least one JSONB field is present (BUG-005
	// response direction).
	needsJSON := hasJSONBProperty(schema)
	needsTypes := hasOpenAPITypesCast(schema)
	needsPgtypex := hasPgtypexColumn(schema, ddlTables, name)
	sb.WriteString("package service\n\nimport (\n")
	if needsJSON {
		sb.WriteString("\t\"encoding/json\"\n\n")
	}
	sb.WriteString("\t" + apiImportLine(modulePath, dg.ApiSuffix) + "\n")
	sb.WriteString("\t\"" + modulePath + "/internal/db\"\n")
	if needsTypes {
		// openapi_types is the import alias oapi-codegen uses in its own
		// generated code for github.com/oapi-codegen/runtime/types
		// (Email, UUID, …). Keep the alias identical so the cast
		// expressions compile against the same named types the api
		// struct fields declare.
		sb.WriteString("\n\topenapi_types \"github.com/oapi-codegen/runtime/types\"\n")
	}
	if needsPgtypex {
		sb.WriteString("\t\"github.com/park-jun-woo/ssac/pkg/pgtypex\"\n")
	}
	sb.WriteString(")\n\n")
	writeConvertFunc(&sb, name, schema, ddlTables, dg.FuncPrefix)

	fileName := fffile.EnsureUnique(fffile.FileNameForFunc(funcName), used)
	return fffile.WriteIfNotPreserved(filepath.Join(serviceDir, fileName), []byte(sb.String()))
}
