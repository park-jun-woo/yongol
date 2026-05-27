import {
  Controller,
  Post,
  Param,
  Req,
} from '@nestjs/common';
import { PauseWorkflowService } from './pauseWorkflow.service';

@Controller('workflows')
export class PauseWorkflowController {
  constructor(private readonly service: PauseWorkflowService) {}

  @Post(':id/pause')
  async pauseWorkflow(
    @Req() req: any,
    @Param() params: any,
  ) {
    return this.service.pauseWorkflow(params, req.user);
  }
}
