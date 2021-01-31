import { Test, TestingModule } from '@nestjs/testing';
import { DanbooruController } from './danbooru.controller';

describe('DanbooruController', () => {
  let controller: DanbooruController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [DanbooruController],
    }).compile();

    controller = module.get<DanbooruController>(DanbooruController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
