import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { SetScheduleService } from './setSchedule.service';

@Controller('workflows')
export class SetScheduleController {
  constructor(private readonly service: SetScheduleService) {}

  @Post(':id/schedule')
  async setSchedule(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.setSchedule(params, body, req.user);
  }
}
