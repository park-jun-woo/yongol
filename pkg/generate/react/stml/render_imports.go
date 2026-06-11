//ff:func feature=stml-gen type=generator control=iteration dimension=1 topic=import-collect
//ff:what importSet을 기반으로 React 임포트 블록 문자열을 생성한다
package stml

import (
	"fmt"
	"strings"
)

// renderImports generates the import block string.
func renderImports(is importSet, opt GenerateOptions) string {
	var lines []string

	if opt.UseClient {
		lines = append(lines, "'use client'\n")
	}
	// react core hooks (useEffect = the Phase004 document.title mount
	// effect; without it the line stays byte-identical to the
	// useState-only output)
	var reactImports []string
	if is.useEffect {
		reactImports = append(reactImports, "useEffect")
	}
	if is.useState {
		reactImports = append(reactImports, "useState")
	}
	if len(reactImports) > 0 {
		lines = append(lines, fmt.Sprintf("import { %s } from 'react'", strings.Join(reactImports, ", ")))
	}

	// tanstack query
	var queryImports []string
	if is.useQuery {
		queryImports = append(queryImports, "useQuery")
	}
	if is.useMutation {
		queryImports = append(queryImports, "useMutation")
	}
	if is.useQueryClient {
		queryImports = append(queryImports, "useQueryClient")
	}
	if len(queryImports) > 0 {
		lines = append(lines, fmt.Sprintf("import { %s } from '@tanstack/react-query'", strings.Join(queryImports, ", ")))
	}

	// react-router
	var routerImports []string
	if is.useLink {
		routerImports = append(routerImports, "Link")
	}
	if is.useParams {
		routerImports = append(routerImports, "useParams")
	}
	if is.useNavigate {
		routerImports = append(routerImports, "useNavigate")
	}
	if is.useOutletCtx {
		routerImports = append(routerImports, "useOutletContext")
	}
	if len(routerImports) > 0 {
		lines = append(lines, fmt.Sprintf("import { %s } from 'react-router-dom'", strings.Join(routerImports, ", ")))
	}

	// react-hook-form
	if is.useForm {
		lines = append(lines, "import { useForm } from 'react-hook-form'")
	}

	// bearer session store (zustand)
	if is.useAuthStore {
		lines = append(lines, "import { useAuthStore } from '@/stores/auth'")
	}

	// zod + zodResolver
	if is.useZod {
		lines = append(lines, "import { z } from 'zod'")
		lines = append(lines, "import { zodResolver } from '@hookform/resolvers/zod'")
	}

	// ui components
	if is.useButton {
		lines = append(lines, "import { Button } from '@/components/ui/Button'")
	}
	if is.useInput {
		lines = append(lines, "import { Input } from '@/components/ui/Input'")
	}
	if is.useTable {
		lines = append(lines, "import { Table, THead, TBody, TR, TH, TD } from '@/components/ui/Table'")
	}

	// api client
	lines = append(lines, fmt.Sprintf("import { api } from '%s'", opt.APIImportPath))

	// components
	for _, comp := range is.components {
		lines = append(lines, fmt.Sprintf("import %s from '@/components/%s'", comp, comp))
	}

	// custom.ts
	if is.customFile != "" {
		lines = append(lines, fmt.Sprintf("import * as custom from './%s'", is.customFile))
	}

	return strings.Join(lines, "\n")
}
