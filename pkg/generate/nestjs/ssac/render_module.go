//ff:func feature=gen-nestjs type=generator control=iteration dimension=2
//ff:what RenderModule — feature 단위 NestJS module TypeScript 소스 생성 (../ 경로)

package ssac

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderModule produces a NestJS module file that imports the controllers
// and services for a given feature. Each ServicePlan contributes one
// controller and one service to the module.
func RenderModule(feature string, plans []*ir.ServicePlan) (string, error) {
	if feature == "" {
		return "", fmt.Errorf("RenderModule: empty feature name")
	}

	var b strings.Builder

	b.WriteString("import { Module } from '@nestjs/common';\n")
	b.WriteString("import { PrismaModule } from '../prisma/prisma.module';\n")

	needsQueue := false
	needsAuthz := false
	// Collect cross-feature @call dependencies.
	callFeatures := make(map[string]bool)
	for _, p := range plans {
		if hasPublishOp(p.Ops) {
			needsQueue = true
		}
		if hasAuthOp(p.Ops) {
			needsAuthz = true
		}
		for _, op := range p.Ops {
			if op.Kind == ir.OpCall && op.Call != nil && op.Call.TargetFeature != "" {
				if op.Call.TargetFeature != strings.ToLower(feature) {
					callFeatures[op.Call.TargetFeature] = true
				}
			}
		}
	}
	if needsQueue {
		b.WriteString("import { QueueModule } from '../queue/queue.module';\n")
	}
	if needsAuthz {
		b.WriteString("import { AuthzModule } from '../authz/authz.module';\n")
	}

	// Import cross-feature modules.
	sortedFeatures := make([]string, 0, len(callFeatures))
	for cf := range callFeatures {
		sortedFeatures = append(sortedFeatures, cf)
	}
	sort.Strings(sortedFeatures)
	for _, cf := range sortedFeatures {
		modName := strings.ToUpper(cf[:1]) + cf[1:] + "Module"
		b.WriteString(fmt.Sprintf("import { %s } from '../%s/%s.module';\n", modName, cf, cf))
	}

	// Import each controller and service
	for _, p := range plans {
		baseName := lcFirst(p.OperationID)
		b.WriteString(fmt.Sprintf("import { %sController } from './%s.controller';\n",
			p.OperationID, baseName))
		b.WriteString(fmt.Sprintf("import { %sService } from './%s.service';\n",
			p.OperationID, baseName))
	}
	b.WriteString("\n")

	// Module decorator
	moduleName := strings.ToUpper(feature[:1]) + feature[1:] + "Module"
	b.WriteString("@Module({\n")

	// Imports
	b.WriteString("  imports: [\n")
	b.WriteString("    PrismaModule,\n")
	if needsQueue {
		b.WriteString("    QueueModule,\n")
	}
	if needsAuthz {
		b.WriteString("    AuthzModule,\n")
	}
	for _, cf := range sortedFeatures {
		modName := strings.ToUpper(cf[:1]) + cf[1:] + "Module"
		b.WriteString(fmt.Sprintf("    %s,\n", modName))
	}
	b.WriteString("  ],\n")

	// Controllers
	b.WriteString("  controllers: [\n")
	for _, p := range plans {
		b.WriteString(fmt.Sprintf("    %sController,\n", p.OperationID))
	}
	b.WriteString("  ],\n")

	// Providers
	b.WriteString("  providers: [\n")
	for _, p := range plans {
		b.WriteString(fmt.Sprintf("    %sService,\n", p.OperationID))
	}
	b.WriteString("  ],\n")
	b.WriteString("})\n")
	b.WriteString(fmt.Sprintf("export class %s {}\n", moduleName))

	return b.String(), nil
}
