import {
  Controller,
  Get,
  Query,
  Req,
} from '@nestjs/common';
import { ListAuditLogsService } from './listAuditLogs.service';

@Controller('audit-logs')
export class ListAuditLogsController {
  constructor(private readonly service: ListAuditLogsService) {}

  @Get('')
  async listAuditLogs(
    @Req() req: any,
    @Query() query: any,
  ) {
    return this.service.listAuditLogs(query, req.user);
  }
}
