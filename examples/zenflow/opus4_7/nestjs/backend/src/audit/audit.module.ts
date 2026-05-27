import { Module } from '@nestjs/common';
import { PrismaModule } from '../prisma/prisma.module';
import { AuthzModule } from '../authz/authz.module';
import { GetAuditLogController } from './getAuditLog.controller';
import { GetAuditLogService } from './getAuditLog.service';
import { GetRecentAuditLogsController } from './getRecentAuditLogs.controller';
import { GetRecentAuditLogsService } from './getRecentAuditLogs.service';
import { ListAuditLogsController } from './listAuditLogs.controller';
import { ListAuditLogsService } from './listAuditLogs.service';

@Module({
  imports: [
    PrismaModule,
    AuthzModule,
  ],
  controllers: [
    GetAuditLogController,
    GetRecentAuditLogsController,
    ListAuditLogsController,
  ],
  providers: [
    GetAuditLogService,
    GetRecentAuditLogsService,
    ListAuditLogsService,
  ],
  exports: [
    GetAuditLogService,
    GetRecentAuditLogsService,
    ListAuditLogsService,
  ],
})
export class AuditModule {}
