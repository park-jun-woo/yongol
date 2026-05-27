import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { ListExecutionLogsService } from './listExecutionLogs.service';

@Controller('workflows')
export class ListExecutionLogsController {
  constructor(private readonly service: ListExecutionLogsService) {}

  @Get(':id/execution-logs')
  async listExecutionLogs(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.listExecutionLogs(params, body, req.user);
  }
}
