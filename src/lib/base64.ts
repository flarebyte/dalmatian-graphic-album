import { TextDecoder, TextEncoder } from 'text-encoding';

const toArray = (base64: string): Uint8Array => {
  const letters = new TextEncoder('utf-8').encode(base64);
  return letters.map(v => v - 65);
};

const fromArray = (uint8array: Uint8Array): string =>
  new TextDecoder('utf-8').decode(uint8array);

export { toArray, fromArray };
