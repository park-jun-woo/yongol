import {
  Controller,
  Post,
  Param,
  Body,
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
    @Body() body: any,
  ) {
    return this.service.executeWithReport(params, body, req.user);
  }
}
