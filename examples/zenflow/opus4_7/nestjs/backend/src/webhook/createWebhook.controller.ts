import {
  Controller,
  Post,
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
    @Body() body: any,
  ) {
    return this.service.createWebhook(body, req.user);
  }
}
