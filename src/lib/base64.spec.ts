// tslint:disable:no-expression-statement
import test from 'ava';

import { toArray } from './base64';

const base64Example = 'ABCghjk456+/';

test('toArray', t => {
  const actual = toArray(base64Example);
  const expected = new Uint8Array([
    0,
    1,
    2,
    32,
    33,
    35,
    36,
    56,
    57,
    58,
    62,
    63
  ]);
  t.is(actual[0], expected[0]);
  t.is(actual[1], expected[1]);
  t.is(actual[2], expected[2]);
  t.is(actual[3], expected[3]);
  t.is(actual[4], expected[4]);
  t.is(actual[5], expected[5]);
});
