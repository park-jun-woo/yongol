//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeModuleImportsSection — @Module imports: [...] 배열 출력

package ssac

import "strings"

// writeModuleImportsSection writes the imports array of the @Module decorator.
func writeModuleImportsSection(b *strings.Builder, deps moduleDeps) {
	b.WriteString("  imports: [\n")
	b.WriteString("    PrismaModule,\n")
	if deps.NeedsQueue {
		b.WriteString("    QueueModule,\n")
	}
	if deps.NeedsAuthz {
		b.WriteString("    AuthzModule,\n")
	}
	writeCrossFeatureModuleRefs(b, deps.CrossFeatures)
	b.WriteString("  ],\n")
}
