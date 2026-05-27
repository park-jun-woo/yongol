import {
  Controller,
  Post,
  Param,
  Req,
} from '@nestjs/common';
import { ExecuteWithReportService } from './executeWithReport.service';

@Controller('workflows')
export class ExecuteWithReportController {
  constructor(private readonly service: ExecuteWithReportService) {}

  @Post(':id/execute-with-report')
  async executeWithReport(
    @Req() req: any,
    @Param() params: any,
  ) {
    return this.service.executeWithReport(params, req.user);
  }
}
