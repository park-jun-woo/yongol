import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { ListAuditLogsService } from './listAuditLogs.service';

@Controller('audit-logs')
export class ListAuditLogsController {
  constructor(private readonly service: ListAuditLogsService) {}

  @Get('')
  async listAuditLogs(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.listAuditLogs(params, body, req.user);
  }
}
