import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { GetExecutionDetailService } from './getExecutionDetail.service';

@Controller('execution-logs')
export class GetExecutionDetailController {
  constructor(private readonly service: GetExecutionDetailService) {}

  @Get(':id/detail')
  async getExecutionDetail(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.getExecutionDetail(params, body, req.user);
  }
}
