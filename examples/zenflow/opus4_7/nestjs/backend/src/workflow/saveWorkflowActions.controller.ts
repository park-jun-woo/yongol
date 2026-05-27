import {
  Controller,
  Put,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { SaveWorkflowActionsService } from './saveWorkflowActions.service';

@Controller('workflows')
export class SaveWorkflowActionsController {
  constructor(private readonly service: SaveWorkflowActionsService) {}

  @Put(':id/actions')
  async saveWorkflowActions(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.saveWorkflowActions(params, body, req.user);
  }
}
