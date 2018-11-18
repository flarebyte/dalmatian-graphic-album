import { flatbuffers } from 'flatbuffers';
import { Image } from './schema/dalmatian_generated';

const readDalmatianImage = (data: Uint8Array): Image =>
  Image.getRootAsImage(new flatbuffers.ByteBuffer(data));

export { readDalmatianImage };
