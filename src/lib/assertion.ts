import { flatbuffers } from 'flatbuffers';
import { AnyURI, Assertion, Bool, Curie, Float, IRI, IsoDateTime, Localized,
  ObjectValue, Text, Uint16, Uint32, Uint8  } from './schema/dalmatian_generated';
import { range} from './utils';

// tslint:disable no-if-statement

interface ResourceIdentifier {
  readonly prefix: string;
  readonly path: string;
}

interface LocalizedText {
  readonly  language: string;
  readonly text: string;
}

interface AnyURIValue {
  readonly  value: string;
  readonly  mediaType: string;
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

const noLocalized = new Localized();
const getAsLocalizedText = (assertion: Assertion | null): LocalizedText | null => {
  const value = assertion ? assertion.objectValue(noLocalized) : null;
  const localizedText = value ? value.text(): null;
  const localizedLanguage = value ? value.language(): null;
  const language = localizedLanguage === null ? 'en-GB' : localizedLanguage;
  return localizedText === null ? null : {  language, text: localizedText};
}

const noAnyURI = new AnyURI();
const getAsAnyURIValue = (assertion: Assertion | null): AnyURIValue | null => {
  const value = assertion ? assertion.objectValue(noAnyURI) : null;
  const valueOrNull = value ? value.value(): null;
  const mediaTypeOrNull = value ? value.mediaType(): null;
  const mediaType = mediaTypeOrNull === null ? 'text/html' : mediaTypeOrNull;
  return valueOrNull === null ? null : {  mediaType, value: valueOrNull};
}

const numberOrZero = (value: number | null) => value === null ? 0 : value;
const noFloat = new Float();
const getAsFloatArray = (assertion: Assertion | null): ReadonlyArray<number> => {
  if (!assertion) { return []; }

  const objValue = assertion.objectValue(noFloat);
  if (!objValue) { return []; }

  const indexes = range(objValue.valueLength());

  const  getFloatValue = (i: number) => numberOrZero(objValue.value(i))

  return indexes.map(getFloatValue);
}

const noUint8 = new Uint8();
const getAsUint8Array= (assertion: Assertion | null): Uint8Array => {
  if (!assertion) { return new Uint8Array([]); }

  const objValue = assertion.objectValue(noUint8);
  if (!objValue) { return new Uint8Array([]); }

  const indexes = range(objValue.valueLength());

  const  getFloatValue = (i: number) => numberOrZero(objValue.value(i))

  return new Uint8Array(indexes.map(getFloatValue));
}

const noUint16 = new Uint16();
const getAsUint16Array= (assertion: Assertion | null): Uint16Array => {
  if (!assertion) { return new Uint16Array([]); }

  const objValue = assertion.objectValue(noUint16);
  if (!objValue) { return new Uint16Array([]); }

  const indexes = range(objValue.valueLength());

  const  getFloatValue = (i: number) => numberOrZero(objValue.value(i))

  return new Uint16Array(indexes.map(getFloatValue));
}

const noUint32 = new Uint32();
const getAsUint32Array= (assertion: Assertion | null): Uint32Array => {
  if (!assertion) { return new Uint32Array([]); }

  const objValue = assertion.objectValue(noUint32);
  if (!objValue) { return new Uint32Array([]); }

  const indexes = range(objValue.valueLength());

  const  getFloatValue = (i: number) => numberOrZero(objValue.value(i))

  return new Uint32Array(indexes.map(getFloatValue));
}


export { ObjectValue, matchPredicate, getAsBoolean, getAsFloatArray,
   getAsLocalizedText, getAsString, getAsIsoDateTimeString, getAsAnyURIValue,
    getAsUint8Array, getAsUint16Array, getAsUint32Array};
