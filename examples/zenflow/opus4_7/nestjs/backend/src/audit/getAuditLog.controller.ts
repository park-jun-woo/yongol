import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { GetAuditLogService } from './getAuditLog.service';

@Controller('audit-logs')
export class GetAuditLogController {
  constructor(private readonly service: GetAuditLogService) {}

  @Get(':id')
  async getAuditLog(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.getAuditLog(params, body, req.user);
  }
}
