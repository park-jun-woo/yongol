import {
  Controller,
  Get,
  Param,
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
  ) {
    return this.service.getExecutionDetail(params, req.user);
  }
}
