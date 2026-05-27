import {
  Controller,
  Post,
  Param,
  Body,
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
    @Body() body: any,
  ) {
    return this.service.activateWorkflow(params, body, req.user);
  }
}
