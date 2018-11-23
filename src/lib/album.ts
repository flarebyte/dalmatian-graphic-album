import { flatbuffers } from 'flatbuffers';
import { Album, IRI, Curie } from './schema/dalmatian_generated';

const fromUint8Array = (data: Uint8Array): Album =>
  Album.getRootAsAlbum(new flatbuffers.ByteBuffer(data));

const orEmptyString = (value: string | null ) => value === null ? '' : value;
const orEmptyCurie = (value: Curie | null ) => value === null ? '' : orEmptyString(value.prefix());

const iriToString = (iri: IRI | null)  => {
  return iri === null ? '' : orEmptyCurie(iri.curie()) + orEmptyString(iri.path());
}

const getValueAsString = (album: Album, index: number) => (subject: string, predicate: string) {
  const length = album.metadataLength();
  const tripleValue = album.metadata(index);
  const subjectIRI = iriToString(tripleValue.subject());
  const predicateIRI = tripleValue.predicate();
  const objectValue = tripleValue.objectValue();

 }
  
export { Album, fromUint8Array };
