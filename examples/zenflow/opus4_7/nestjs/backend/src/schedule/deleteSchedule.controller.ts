import {
  Controller,
  Delete,
  Param,
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
  ) {
    return this.service.deleteSchedule(params, req.user);
  }
}
