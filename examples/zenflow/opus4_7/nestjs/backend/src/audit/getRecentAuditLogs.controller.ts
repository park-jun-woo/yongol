import {
  Controller,
  Get,
  Req,
} from '@nestjs/common';
import { GetRecentAuditLogsService } from './getRecentAuditLogs.service';

@Controller('audit-logs')
export class GetRecentAuditLogsController {
  constructor(private readonly service: GetRecentAuditLogsService) {}

  @Get('recent')
  async getRecentAuditLogs(
    @Req() req: any,
  ) {
    return this.service.getRecentAuditLogs(req.user);
  }
}
