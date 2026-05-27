import {
  Controller,
  Post,
  Param,
  Req,
} from '@nestjs/common';
import { CreateWorkflowVersionService } from './createWorkflowVersion.service';

@Controller('workflows')
export class CreateWorkflowVersionController {
  constructor(private readonly service: CreateWorkflowVersionService) {}

  @Post(':id/new-version')
  async createWorkflowVersion(
    @Req() req: any,
    @Param() params: any,
  ) {
    return this.service.createWorkflowVersion(params, req.user);
  }
}
