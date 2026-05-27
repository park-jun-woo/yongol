import {
  Controller,
  Post,
  Param,
  Body,
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
    @Body() body: any,
  ) {
    return this.service.pauseWorkflow(params, body, req.user);
  }
}
