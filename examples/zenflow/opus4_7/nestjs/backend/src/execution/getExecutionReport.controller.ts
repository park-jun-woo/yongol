import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { GetExecutionReportService } from './getExecutionReport.service';

@Controller('execution-logs')
export class GetExecutionReportController {
  constructor(private readonly service: GetExecutionReportService) {}

  @Get(':id/report')
  async getExecutionReport(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.getExecutionReport(params, body, req.user);
  }
}
