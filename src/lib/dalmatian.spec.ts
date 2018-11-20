// tslint:disable:no-expression-statement
import test from 'ava';

import { readDalmatianImage } from './dalmatian';
import { Color } from './schema/dalmatian_generated';

test('toArray', t => {
  const data = new Uint8Array();
  const actual = readDalmatianImage(data);
   const publishing = actual.publishing(0);
   
  t.is(publishing., );
  
});
