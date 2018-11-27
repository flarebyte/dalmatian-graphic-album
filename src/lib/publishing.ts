import { Publishing } from './schema/dalmatian_generated';

const getMetadata = (publishing: Publishing | null) =>
  publishing ? publishing.metadata() : null;

export { getMetadata };
