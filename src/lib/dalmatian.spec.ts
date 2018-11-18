// tslint:disable:no-expression-statement
import test from 'ava';

import { readDalmatianImage } from './dalmatian';

test('toArray', t => {
  const data = new Uint8Array();
  const actual = readDalmatianImage(data);
  t.is(actual.color.length, 0);
  t.is(actual.iriSpace.length, 0);
  t.is(actual.processor.length, 0);
  t.is(actual.publishing.length, 0);
  t.is(actual.layer.length, 0);
});
