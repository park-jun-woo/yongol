import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { ArchiveWorkflowService } from './archiveWorkflow.service';

@Controller('workflows')
export class ArchiveWorkflowController {
  constructor(private readonly service: ArchiveWorkflowService) {}

  @Post(':id/archive')
  async archiveWorkflow(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.archiveWorkflow(params, body, req.user);
  }
}
