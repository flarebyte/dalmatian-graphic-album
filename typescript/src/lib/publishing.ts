import { Publishing } from './dalmatian_generated';

const getMetadata = (publishing: Publishing | null) =>
  publishing ? publishing.metadata() : null;

export { getMetadata };
