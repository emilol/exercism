// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */
export function timeToMixJuice(name) {
  switch (name) {
    case 'Pure Strawberry Joy':
      return 0.5;
    case 'Energizer':
    case 'Green Garden':
      return 1.5;
    case 'Tropical Island':
      return 3;
    case 'All or Nothing':
      return 5;
    default:
      return 2.5;
  }
}

function wedgeCountFor(lime) {
  switch (lime) {
    case 'small':
      return 6;
    case 'medium':
      return 8;
    case 'large':
      return 10;
    default:
      throw new Error('Unknown lime');
  }
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */
export function limesToCut(wedgesNeeded, limes) {
  let wedgeCount = 0;
  let limeCount = 0;
  while (wedgesNeeded > wedgeCount && limes.length > 0) {
    const lime = limes.shift();
    wedgeCount += wedgeCountFor(lime);
    limeCount++;
  }
  return limeCount;
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
export function remainingOrders(timeLeft, orders) {
  let remainingTime = timeLeft;
  do {
    let order = orders.shift();
    let time = timeToMixJuice(order);
    remainingTime -= time;
  } while (remainingTime > 0 && orders.length > 0);
  return orders;
}
