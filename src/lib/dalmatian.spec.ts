// tslint:disable:no-expression-statement
import test from 'ava';

import { readDalmatianImage } from './dalmatian';
import { Color } from './schema/dalmatian_generated';

test('toArray', t => {
  const data = new Uint8Array();
  const actual = readDalmatianImage(data);
   const color = actual.color(0);
   
  t.is(color., );
  
});
