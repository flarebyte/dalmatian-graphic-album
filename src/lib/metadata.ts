import { flatbuffers } from 'flatbuffers';
import { Metadata } from './schema/dalmatian_generated';

const getValueAsString = (metadata: Metadata, index: number) => (subject: string, predicate: string) {
  const tripleValue =metadata.assertion(0);
  const subjectIRI = iriToString(tripleValue.subject());
  const predicateIRI = tripleValue.predicate();
  const objectValue = tripleValue.objectValue();

 }
  
export { Album, fromUint8Array };
