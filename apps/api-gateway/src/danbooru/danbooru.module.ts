import { Module } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { DANBOORU_V1_PACKAGE_NAME } from 'src/generated/apis/danbooru';
import { DanbooruController } from './danbooru.controller';
import { DanbooruService } from './danbooru.service';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: DANBOORU_V1_PACKAGE_NAME,
        inject: [ConfigService],
        useFactory: async (configService: ConfigService<NodeJS.ProcessEnv>) => ({
          name: DANBOORU_V1_PACKAGE_NAME,
          transport: Transport.GRPC,
          options: {
            package: DANBOORU_V1_PACKAGE_NAME,
            protoPath: join(__dirname, '../../apis/danbooru.proto'),
            url: configService.get<string>('SVC_DANBOORU_GRPC_URL'),
          },
        }),
      },
    ]),
  ],
  controllers: [DanbooruController],
  providers: [DanbooruService],
})
export class DanbooruModule {}
