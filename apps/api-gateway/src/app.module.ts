import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { KonachanModule } from './konachan/konachan.module';

@Module({
  imports: [KonachanModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
