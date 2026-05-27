import {
  Controller,
  Post,
  Param,
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
  ) {
    return this.service.cloneTemplate(params, req.user);
  }
}
