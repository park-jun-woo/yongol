import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { AutoAssignWorkflowService } from './autoAssignWorkflow.service';

@Controller('workflows')
export class AutoAssignWorkflowController {
  constructor(private readonly service: AutoAssignWorkflowService) {}

  @Post(':id/auto-assign')
  async autoAssignWorkflow(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.autoAssignWorkflow(params, body, req.user);
  }
}
