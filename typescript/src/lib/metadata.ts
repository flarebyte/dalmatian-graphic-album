import { matchPredicate, ResourceIdentifier } from './assertion';
import { Assertion, Metadata } from './dalmatian_generated';
import { range } from './utils';

const getId = (metadata: Metadata | null): number =>
  metadata === null ? -1 : metadata.id();

const notNullAssertion = (assertion: Assertion | null) => assertion !== null;
const voidAssertion = new Assertion();

const discardNullAssertions = (
  assertions: ReadonlyArray<Assertion | null>
): ReadonlyArray<Assertion> =>
  assertions
    .filter(notNullAssertion)
    .map(a => (a === null ? voidAssertion : a));

const findAssertions = (iri: ResourceIdentifier) => (
  metadata: Metadata | null
): ReadonlyArray<Assertion> => {
  const length = metadata ? metadata.assertionLength() : 0;
  const getAssertion = (i: number) => (metadata ? metadata.assertion(i) : null);
  const indexes = range(length);
  const matchIri = matchPredicate(iri);
  const keepAssertion = (i: number) => {
    const assertion = getAssertion(i);
    const isMatched = matchIri(assertion);
    return isMatched ? assertion : null;
  };

  return discardNullAssertions(indexes.map(keepAssertion));
};

const findFirstAssertion = (iri: ResourceIdentifier) => (
  metadata: Metadata | null
): Assertion | null => findAssertions(iri)(metadata)[0];

export { getId, findAssertions, findFirstAssertion };
