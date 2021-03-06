import {
  AnyURI,
  Assertion,
  Bool,
  Curie,
  Float,
  IRI,
  IsoDateTime,
  Localized,
  ObjectValue,
  Text,
  Uint16,
  Uint32,
  Uint8,
  Vector2d,
  Vector2dFloat,
  Vector3d,
  Vector3dFloat,
  VectorColorRGBA
} from './schema/dalmatian_generated';
import { range } from './utils';

// tslint:disable no-if-statement

interface ResourceIdentifier {
  readonly prefix: string;
  readonly path: string;
}

interface LocalizedText {
  readonly language: string;
  readonly text: string;
}

interface AnyURIValue {
  readonly value: string;
  readonly mediaType: string;
}

interface VectorXY {
  readonly x: number;
  readonly y: number;
}

interface VectorXYZ {
  readonly x: number;
  readonly y: number;
  readonly z: number;
}

interface VectorXYFloat {
  readonly x: number;
  readonly y: number;
}

interface VectorXYZFloat {
  readonly x: number;
  readonly y: number;
  readonly z: number;
}

interface VectorRGBA {
  readonly r: number;
  readonly g: number;
  readonly b: number;
  readonly a: number;
}

const matchString = (expected: string) => (actual: string | null) => {
  return actual ? expected === actual : false;
};

const matchCurie = (expected: string) => (actual: Curie | null) => {
  const prefix = actual ? actual.prefix() : null;
  return prefix ? matchString(expected)(prefix) : false;
};

const matchIRI = (expected: ResourceIdentifier) => (actual: IRI | null) => {
  const curie = actual ? actual.curie() : null;
  return matchCurie(expected.prefix)(curie);
};

const matchPredicate = (expected: ResourceIdentifier) => (
  actual: Assertion | null
) => {
  const predicate = actual ? actual.predicate() : null;
  return matchIRI(expected)(predicate);
};

const noText = new Text();
const getAsString = (assertion: Assertion): string | null => {
  const value = assertion.objectValue(noText);
  return value ? value.value() : null;
};

const noBool = new Bool();
const getAsBoolean = (assertion: Assertion): boolean | null => {
  const value = assertion.objectValue(noBool);
  return value ? value.value() : null;
};

const noIsoDateTime = new IsoDateTime();
const getAsIsoDateTimeString = (assertion: Assertion): string | null => {
  const value = assertion.objectValue(noIsoDateTime);
  return value ? value.value() : null;
};

const noLocalized = new Localized();
const getAsLocalizedText = (
  assertion: Assertion | null
): LocalizedText | null => {
  const value = assertion ? assertion.objectValue(noLocalized) : null;
  const localizedText = value ? value.text() : null;
  const localizedLanguage = value ? value.language() : null;
  const language = localizedLanguage === null ? 'en-GB' : localizedLanguage;
  return localizedText === null ? null : { language, text: localizedText };
};

const noAnyURI = new AnyURI();
const getAsAnyURIValue = (assertion: Assertion): AnyURIValue | null => {
  const value = assertion.objectValue(noAnyURI);
  const valueOrNull = value ? value.value() : null;
  const mediaTypeOrNull = value ? value.mediaType() : null;
  const mediaType = mediaTypeOrNull === null ? 'text/html' : mediaTypeOrNull;
  return valueOrNull === null ? null : { mediaType, value: valueOrNull };
};

const noIRI = new IRI();
const getAsIRI = (assertion: Assertion): ResourceIdentifier | null => {
  const value = assertion.objectValue(noIRI);
  const curieOrNull = value ? value.curie() : null;
  const prefixOrNull = curieOrNull ? curieOrNull.prefix() : null;
  const pathOrNull = value ? value.path() : null;
  const prefixOrEmpty = prefixOrNull ? prefixOrNull : '';
  return pathOrNull === null
    ? null
    : { prefix: prefixOrEmpty, path: pathOrNull };
};

const noVector2d = new Vector2d();
const getAsVectorXY = (assertion: Assertion): VectorXY | null => {
  const value = assertion.objectValue(noVector2d);
  const v = value ? value.value() : null;
  return v === null ? null : { x: v.x(), y: v.y() };
};

const noVector3d = new Vector3d();
const getAsVectorXYZ = (assertion: Assertion): VectorXYZ | null => {
  const value = assertion.objectValue(noVector3d);
  const v = value ? value.value() : null;
  return v === null ? null : { x: v.x(), y: v.y(), z: v.z() };
};

const noVector2dFloat = new Vector2dFloat();
const getAsVectorXYFloat = (assertion: Assertion): VectorXYFloat | null => {
  const value = assertion.objectValue(noVector2dFloat);
  const v = value ? value.value() : null;
  return v === null ? null : { x: v.x(), y: v.y() };
};

const noVector3dFloat = new Vector3dFloat();
const getAsVectorXYZFloat = (assertion: Assertion): VectorXYZFloat | null => {
  const value = assertion.objectValue(noVector3dFloat);
  const v = value ? value.value() : null;
  return v === null ? null : { x: v.x(), y: v.y(), z: v.z() };
};

const noVectorColorRGBA = new VectorColorRGBA();
const getAsVectorRGBA = (assertion: Assertion): VectorRGBA | null => {
  const value = assertion.objectValue(noVectorColorRGBA);
  const v = value ? value.value() : null;
  return v === null ? null : { r: v.r(), g: v.g(), b: v.b(), a: v.a() };
};

const numberOrZero = (value: number | null) => (value === null ? 0 : value);
const noFloat = new Float();
const getAsFloatArray = (assertion: Assertion): ReadonlyArray<number> => {
  const objValue = assertion.objectValue(noFloat);
  if (!objValue) {
    return [];
  }

  const indexes = range(objValue.valueLength());

  const getFloatValue = (i: number) => numberOrZero(objValue.value(i));

  return indexes.map(getFloatValue);
};

const noUint8 = new Uint8();
const getAsUint8Array = (assertion: Assertion): Uint8Array => {
  const objValue = assertion.objectValue(noUint8);
  if (!objValue) {
    return new Uint8Array([]);
  }

  const indexes = range(objValue.valueLength());

  const getFloatValue = (i: number) => numberOrZero(objValue.value(i));

  return new Uint8Array(indexes.map(getFloatValue));
};

const noUint16 = new Uint16();
const getAsUint16Array = (assertion: Assertion): Uint16Array => {
  const objValue = assertion.objectValue(noUint16);
  if (!objValue) {
    return new Uint16Array([]);
  }

  const indexes = range(objValue.valueLength());

  const getFloatValue = (i: number) => numberOrZero(objValue.value(i));

  return new Uint16Array(indexes.map(getFloatValue));
};

const noUint32 = new Uint32();
const getAsUint32Array = (assertion: Assertion): Uint32Array => {
  const objValue = assertion.objectValue(noUint32);
  if (!objValue) {
    return new Uint32Array([]);
  }

  const indexes = range(objValue.valueLength());

  const getFloatValue = (i: number) => numberOrZero(objValue.value(i));

  return new Uint32Array(indexes.map(getFloatValue));
};

export {
  ObjectValue,
  matchPredicate,
  getAsBoolean,
  getAsFloatArray,
  getAsLocalizedText,
  getAsString,
  getAsIsoDateTimeString,
  getAsAnyURIValue,
  getAsIRI,
  getAsUint8Array,
  getAsUint16Array,
  getAsUint32Array,
  getAsVectorXY,
  getAsVectorXYZ,
  getAsVectorXYFloat,
  getAsVectorXYZFloat,
  getAsVectorRGBA,
  ResourceIdentifier
};
