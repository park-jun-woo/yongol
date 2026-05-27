import { Module } from '@nestjs/common';
import { PrismaModule } from '../prisma/prisma.module';
import { AuthzModule } from '../authz/authz.module';
import { DashboardModule } from '../dashboard/dashboard.module';
import { GetExecutionDetailController } from './getExecutionDetail.controller';
import { GetExecutionDetailService } from './getExecutionDetail.service';
import { GetExecutionReportController } from './getExecutionReport.controller';
import { GetExecutionReportService } from './getExecutionReport.service';

@Module({
  imports: [
    PrismaModule,
    AuthzModule,
    DashboardModule,
  ],
  controllers: [
    GetExecutionDetailController,
    GetExecutionReportController,
  ],
  providers: [
    GetExecutionDetailService,
    GetExecutionReportService,
  ],
  exports: [
    GetExecutionDetailService,
    GetExecutionReportService,
  ],
})
export class ExecutionModule {}
