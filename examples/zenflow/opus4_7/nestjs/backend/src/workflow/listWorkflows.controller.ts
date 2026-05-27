import {
  Controller,
  Get,
  Req,
} from '@nestjs/common';
import { ListWorkflowsService } from './listWorkflows.service';

@Controller('workflows')
export class ListWorkflowsController {
  constructor(private readonly service: ListWorkflowsService) {}

  @Get('')
  async listWorkflows(
    @Req() req: any,
  ) {
    return this.service.listWorkflows(req.user);
  }
}
