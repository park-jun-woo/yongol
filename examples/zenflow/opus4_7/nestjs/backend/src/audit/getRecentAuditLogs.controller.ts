import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { GetRecentAuditLogsService } from './getRecentAuditLogs.service';

@Controller('audit-logs')
export class GetRecentAuditLogsController {
  constructor(private readonly service: GetRecentAuditLogsService) {}

  @Get('recent')
  async getRecentAuditLogs(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.getRecentAuditLogs(params, body, req.user);
  }
}
