import { Module } from '@nestjs/common';
import { PrismaModule } from '../../prisma/prisma.module';
import { GetExecutionDetailController } from './getExecutionDetail.controller';
import { GetExecutionDetailService } from './getExecutionDetail.service';
import { GetExecutionReportController } from './getExecutionReport.controller';
import { GetExecutionReportService } from './getExecutionReport.service';

@Module({
  imports: [
    PrismaModule,
  ],
  controllers: [
    GetExecutionDetailController,
    GetExecutionReportController,
  ],
  providers: [
    GetExecutionDetailService,
    GetExecutionReportService,
  ],
})
export class ExecutionModule {}
