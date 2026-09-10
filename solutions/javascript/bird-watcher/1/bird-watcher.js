// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  let birdCount = 0;
  for (let index = 0; index < birdsPerDay.length; index++) {
    const birds = birdsPerDay[index];
    birdCount += birds;
  }

  return birdCount;
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  let birdCount = 0;
  let start = (week - 1) * 7;
  for (let index = start; index < start + 7; index++) {
    const birds = birdsPerDay[index];
    birdCount += birds;
  }
  return birdCount;
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {void} should not return anything
 */
export function fixBirdCountLog(birdsPerDay) {
  for (let index = 0; index < birdsPerDay.length; index += 2) {
    let birds = birdsPerDay[index];
    birdsPerDay[index] = ++birds;
  }
}
