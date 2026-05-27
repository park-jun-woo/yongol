import { Module } from '@nestjs/common';
import { PrismaModule } from '../prisma/prisma.module';
import { AuthzModule } from '../authz/authz.module';
import { CloneTemplateController } from './cloneTemplate.controller';
import { CloneTemplateService } from './cloneTemplate.service';
import { GetTemplateController } from './getTemplate.controller';
import { GetTemplateService } from './getTemplate.service';
import { ListTemplatesController } from './listTemplates.controller';
import { ListTemplatesService } from './listTemplates.service';
import { PublishTemplateController } from './publishTemplate.controller';
import { PublishTemplateService } from './publishTemplate.service';

@Module({
  imports: [
    PrismaModule,
    AuthzModule,
  ],
  controllers: [
    CloneTemplateController,
    GetTemplateController,
    ListTemplatesController,
    PublishTemplateController,
  ],
  providers: [
    CloneTemplateService,
    GetTemplateService,
    ListTemplatesService,
    PublishTemplateService,
  ],
})
export class TemplateModule {}
