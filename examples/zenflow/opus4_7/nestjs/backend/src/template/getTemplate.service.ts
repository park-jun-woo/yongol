import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';

@Injectable()
export class GetTemplateService {
  constructor(
    private readonly prisma: PrismaService,
  ) {}

  async getTemplate(params: any, user?: any): Promise<any> {
    const template = await this.prisma.template.findUnique({ where: { id: params.id } });
    if (!template) {
      throw new HttpException('Template not found', HttpStatus.NOT_FOUND);
    }
    return {
      template: template,
    };
  }
}
