import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { CreateWebhookService } from './createWebhook.service';

@Controller('webhooks')
export class CreateWebhookController {
  constructor(private readonly service: CreateWebhookService) {}

  @Post('')
  async createWebhook(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.createWebhook(params, body, req.user);
  }
}
