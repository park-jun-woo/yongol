import {
  Controller,
  Get,
  Query,
  Req,
} from '@nestjs/common';
import { ListTemplatesService } from './listTemplates.service';

@Controller('templates')
export class ListTemplatesController {
  constructor(private readonly service: ListTemplatesService) {}

  @Get('')
  async listTemplates(
    @Req() req: any,
    @Query() query: any,
  ) {
    return this.service.listTemplates(query, req.user);
  }
}
