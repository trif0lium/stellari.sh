import { Module } from '@nestjs/common';
import { NhentaiService } from './nhentai.service';
import { NhentaiController } from './nhentai.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { NHENTAI_V1_PACKAGE_NAME } from 'src/generated/apis/nhentai';
import { ConfigService } from '@nestjs/config';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: NHENTAI_V1_PACKAGE_NAME,
        inject: [ConfigService],
        useFactory: async (configService: ConfigService<NodeJS.ProcessEnv>) => ({
          name: NHENTAI_V1_PACKAGE_NAME,
          transport: Transport.GRPC,
          options: {
            package: NHENTAI_V1_PACKAGE_NAME,
            protoPath: join(__dirname, '../../apis/nhentai.proto'),
            url: configService.get<string>('SVC_NHENTAI_GRPC_URL'),
          },
        }),
      },
    ]),
  ],
  providers: [NhentaiService],
  controllers: [NhentaiController],
})
export class NhentaiModule {}
