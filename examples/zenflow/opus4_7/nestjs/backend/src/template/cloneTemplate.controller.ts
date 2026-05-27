import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { CloneTemplateService } from './cloneTemplate.service';

@Controller('templates')
export class CloneTemplateController {
  constructor(private readonly service: CloneTemplateService) {}

  @Post(':id/clone')
  async cloneTemplate(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.cloneTemplate(params, body, req.user);
  }
}
