import { Test, TestingModule } from '@nestjs/testing';
import { NhentaiController } from './nhentai.controller';

describe('NhentaiController', () => {
  let controller: NhentaiController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [NhentaiController],
    }).compile();

    controller = module.get<NhentaiController>(NhentaiController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
