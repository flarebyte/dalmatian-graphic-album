// tslint:disable no-if-statement no-let readonly-array no-expression-statement
const maxBufferSize = 10000;

const generateRange = () => {
  const results: number[] = new Array(maxBufferSize);
  let i = 0;
  for (i = 0; i < maxBufferSize; i++) {
    results.push(i);
  }
  return results;
};
const cachedRange: ReadonlyArray<number> = generateRange();

const range = (length: number) => {
  const maxLength = length < maxBufferSize ? length : maxBufferSize;
  return cachedRange.slice(0, maxLength - 1);
};

export { range };
