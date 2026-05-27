import { Global, Module } from '@nestjs/common';
import { ReportService } from './report.service';

@Global()
@Module({
  providers: [ReportService],
  exports: [ReportService],
})
export class ReportModule {}
