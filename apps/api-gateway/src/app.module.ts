import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { KonachanModule } from './konachan/konachan.module';
import { resolve } from 'path';
import { DanbooruModule } from './danbooru/danbooru.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: [resolve(__dirname, '../.env')],
      isGlobal: true,
    }),
    KonachanModule,
    DanbooruModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
