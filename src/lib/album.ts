import { flatbuffers } from 'flatbuffers';
import { getAsIRI, ResourceIdentifier } from './assertion';
import { findAssertions } from './metadata';
import { getMetadata as getPublishingMeta } from './publishing';
import {
  Album,
  Assertion,
  Metadata,
  Publishing
} from './schema/dalmatian_generated';
import { dctermsFormat } from './term';
import { range } from './utils';

const fromUint8Array = (data: Uint8Array): Album =>
  Album.getRootAsAlbum(new flatbuffers.ByteBuffer(data));

const getMetadata = (album: Album) => album.metadata();

const findPublishingByFormat = (album: Album) => (
  format: ResourceIdentifier
): Publishing | null => {
  const length = album.publishingLength();
  const indexes = range(length);
  const getPublishing = (i: number) => (album ? album.publishing(i) : null);
  const getMeta = (pub: Publishing | null) => getPublishingMeta(pub);
  const matchFormatPredicate = (meta: Metadata | null) =>
    meta ? findAssertions(dctermsFormat)(meta) : null;
  const matchFormatValue = (assertion: Assertion | null) =>
    assertion ? format === getAsIRI(assertion) : false;
  const correctFormat = (pub: Publishing | null) => {
    const meta = getMeta(pub);
    const assertions = matchFormatPredicate(meta);
    return assertions ? assertions.filter(matchFormatValue).length > 0 : false;
  };
  return indexes.map(getPublishing).filter(correctFormat)[0];
};

export { Album, fromUint8Array, getMetadata, findPublishingByFormat };
