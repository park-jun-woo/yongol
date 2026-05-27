import {
  Controller,
  Get,
  Param,
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
  ) {
    return this.service.getSchedule(params, req.user);
  }
}
