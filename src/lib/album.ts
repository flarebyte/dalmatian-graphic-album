import { flatbuffers } from 'flatbuffers';
import { Album } from './schema/dalmatian_generated';

const fromUint8Array = (data: Uint8Array): Album =>
  Album.getRootAsAlbum(new flatbuffers.ByteBuffer(data));

export { Album, fromUint8Array };
