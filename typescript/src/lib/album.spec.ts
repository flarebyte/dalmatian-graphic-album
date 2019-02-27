// tslint:disable:no-expression-statement
import test from 'ava';

import { fromUint8Array } from './album';

test('fromUint8Array', t => {
  const data = new Uint8Array();
  const actual = fromUint8Array(data);
  t.is(typeof actual, 'object');
});
