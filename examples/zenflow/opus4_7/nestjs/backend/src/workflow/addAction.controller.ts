import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { AddActionService } from './addAction.service';

@Controller('workflows')
export class AddActionController {
  constructor(private readonly service: AddActionService) {}

  @Post(':id/actions')
  async addAction(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.addAction(params, body, req.user);
  }
}
