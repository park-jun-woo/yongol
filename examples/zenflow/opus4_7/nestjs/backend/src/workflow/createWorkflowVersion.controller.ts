import {
  Controller,
  Post,
  Param,
  Body,
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
    @Body() body: any,
  ) {
    return this.service.createWorkflowVersion(params, body, req.user);
  }
}
