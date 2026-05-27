import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { GetScheduleService } from './getSchedule.service';

@Controller('workflows')
export class GetScheduleController {
  constructor(private readonly service: GetScheduleService) {}

  @Get(':id/schedule')
  async getSchedule(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.getSchedule(params, body, req.user);
  }
}
