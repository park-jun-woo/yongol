import {
  Controller,
  Delete,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { DeleteScheduleService } from './deleteSchedule.service';

@Controller('workflows')
export class DeleteScheduleController {
  constructor(private readonly service: DeleteScheduleService) {}

  @Delete(':id/schedule')
  async deleteSchedule(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.deleteSchedule(params, body, req.user);
  }
}
