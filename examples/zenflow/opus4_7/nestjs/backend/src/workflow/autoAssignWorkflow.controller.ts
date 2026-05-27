import {
  Controller,
  Post,
  Param,
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
  ) {
    return this.service.autoAssignWorkflow(params, req.user);
  }
}
