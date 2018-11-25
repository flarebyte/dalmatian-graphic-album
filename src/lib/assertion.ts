import { flatbuffers } from 'flatbuffers';
import { Assertion, Bool, Curie, Float, IRI, IsoDateTime, ObjectValue, Text  } from './schema/dalmatian_generated';
import { range} from './utils';

// tslint:disable no-if-statement

interface ResourceIdentifier {
  readonly prefix: string;
  readonly path: string;
}

const matchString = (expected: string) => (actual: string | null) => {
  return actual ?  expected === actual : false ;
}

const matchCurie = (expected: string) => (actual: Curie | null) => {
  const prefix = actual ? actual.prefix() : null;
  return prefix ? matchString(expected)(prefix) : false; 
}

const matchIRI = (expected: ResourceIdentifier) => (actual: IRI | null) => {
  const curie = actual ? actual.curie() : null;
  return matchCurie(expected.prefix)(curie);
}

const matchPredicate = (expected: ResourceIdentifier) => (actual: Assertion | null) => {
  const predicate = actual ? actual.predicate() : null;
  return matchIRI(expected)(predicate);
}

const noText = new Text();
const getAsString = (assertion: Assertion | null): string | null => {
  const value = assertion ? assertion.objectValue(noText) : null;
  return value ? value.value() : null;
}

 enum ObjectValue2{
  NONE= 0,
  IRI= 1,
  Localized= 2,
  Text= 3,
  Bool= 4,
  Uint8= 5,
  Uint16= 6,
  Uint32= 7,
  Float= 8,
  Vector2d= 9,
  Vector3d= 10,
  Vector2dFloat= 11,
  Vector3dFloat= 12,
  AnyURI= 13,
  IsoDateTime= 14,
  VectorColorRGBA= 15
};

const noBool = new Bool();
const getAsBoolean = (assertion: Assertion | null): boolean | null => {
  const value = assertion ? assertion.objectValue(noBool) : null;
  return value ? value.value() : null;
}

const noIsoDateTime = new IsoDateTime();
const getAsIsoDateTimeString = (assertion: Assertion | null): string | null => {
  const value = assertion ? assertion.objectValue(noIsoDateTime) : null;
  return value ? value.value() : null;
}

const floatOrZero = (value: number | null) => value === null ? 0 : value;
const noFloat = new Float();
const getAsFloatArray = (assertion: Assertion | null): ReadonlyArray<number> => {
  if (!assertion) { return []; }

  const objValue = assertion.objectValue(noFloat);
  if (!objValue) { return []; }

  const indexes = range(objValue.valueLength());

  const  getFloatValue = (i: number) => floatOrZero(objValue.value(i))

  return indexes.map(getFloatValue);
}


export { ObjectValue, matchPredicate};
