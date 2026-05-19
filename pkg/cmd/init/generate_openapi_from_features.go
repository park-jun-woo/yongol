//ff:func feature=cli-init type=generator control=iteration dimension=1
//ff:what generateOpenAPIFromFeatures — builds OpenAPI YAML with path+operationId stubs from features

package cliinit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// generateOpenAPIFromFeatures builds a complete OpenAPI 3.0.3 document with
// path stubs derived from the features list and writes it to
// <targetDir>/specs/api/openapi.yaml.
func generateOpenAPIFromFeatures(targetDir string, data templateData, feats []features.Feature) error {
	var b strings.Builder

	// Header
	b.WriteString("openapi: 3.0.3\n")
	b.WriteString("info:\n")
	fmt.Fprintf(&b, "  title: %s\n", data.ProjectID)
	fmt.Fprintf(&b, "  description: %s\n", data.Description)
	b.WriteString("  version: 0.1.0\n")
	b.WriteString("servers:\n")
	b.WriteString("  - url: http://localhost:8080\n")
	b.WriteString("security:\n")
	b.WriteString("  - bearerAuth: []\n")
	b.WriteString("components:\n")
	b.WriteString("  securitySchemes:\n")
	b.WriteString("    bearerAuth:\n")
	b.WriteString("      type: http\n")
	b.WriteString("      scheme: bearer\n")
	b.WriteString("      bearerFormat: JWT\n")
	b.WriteString("  schemas:\n")
	b.WriteString("    ErrorResponse:\n")
	b.WriteString("      type: object\n")
	b.WriteString("      required: [error]\n")
	b.WriteString("      properties:\n")
	b.WriteString("        error: { type: string }\n")
	b.WriteString("        code: { type: string }\n")
	b.WriteString("        details:\n")
	b.WriteString("          type: object\n")
	b.WriteString("          additionalProperties: true\n")

	// Group features by URI path
	type opEntry struct {
		Method string
		Op     string
	}
	pathOrder := []string{}
	pathOps := map[string][]opEntry{}

	for _, f := range feats {
		route, err := parseHTTPPath(f.Path)
		if err != nil {
			return err
		}
		if _, exists := pathOps[route.URI]; !exists {
			pathOrder = append(pathOrder, route.URI)
		}
		pathOps[route.URI] = append(pathOps[route.URI], opEntry{
			Method: route.Method,
			Op:     f.Op,
		})
	}

	b.WriteString("paths:\n")
	for _, uri := range pathOrder {
		fmt.Fprintf(&b, "  %s:\n", uri)
		for _, op := range pathOps[uri] {
			fmt.Fprintf(&b, "    %s:\n", op.Method)
			fmt.Fprintf(&b, "      operationId: %s\n", op.Op)
			b.WriteString("      # TODO: add requestBody and responses\n")
			b.WriteString("      responses:\n")
			b.WriteString("        '200':\n")
			b.WriteString("          description: OK\n")
		}
	}

	dest := filepath.Join(targetDir, "specs", "api", "openapi.yaml")
	if err := os.WriteFile(dest, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("write openapi: %w", err)
	}
	return nil
}
