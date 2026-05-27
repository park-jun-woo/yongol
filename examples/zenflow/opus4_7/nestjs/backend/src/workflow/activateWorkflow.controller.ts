import {
  Controller,
  Post,
  Param,
  Req,
} from '@nestjs/common';
import { ActivateWorkflowService } from './activateWorkflow.service';

@Controller('workflows')
export class ActivateWorkflowController {
  constructor(private readonly service: ActivateWorkflowService) {}

  @Post(':id/activate')
  async activateWorkflow(
    @Req() req: any,
    @Param() params: any,
  ) {
    return this.service.activateWorkflow(params, req.user);
  }
}
