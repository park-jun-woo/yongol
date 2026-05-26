//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writePrismaModule — NestJS PrismaModule + PrismaService 파일 기록

package nestjs

import (
	"os"
	"path/filepath"
)

// writePrismaModule writes the PrismaModule and PrismaService files for DI.
func writePrismaModule(srcDir string) error {
	prismaDir := filepath.Join(srcDir, "prisma")
	if err := os.MkdirAll(prismaDir, 0o755); err != nil {
		return err
	}

	svc := `import { Injectable, OnModuleInit, OnModuleDestroy } from '@nestjs/common';
import { PrismaClient } from '@prisma/client';

@Injectable()
export class PrismaService extends PrismaClient implements OnModuleInit, OnModuleDestroy {
  async onModuleInit() {
    await this.$connect();
  }

  async onModuleDestroy() {
    await this.$disconnect();
  }
}
`
	if err := os.WriteFile(filepath.Join(prismaDir, "prisma.service.ts"), []byte(svc), 0o644); err != nil {
		return err
	}

	mod := `import { Global, Module } from '@nestjs/common';
import { PrismaService } from './prisma.service';

@Global()
@Module({
  providers: [PrismaService],
  exports: [PrismaService],
})
export class PrismaModule {}
`
	return os.WriteFile(filepath.Join(prismaDir, "prisma.module.ts"), []byte(mod), 0o644)
}
