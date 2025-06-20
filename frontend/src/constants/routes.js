// src/constants/routes.js
export const GERMAN_ROUTES = Object.freeze({
  GERMAN_ALL_WORDS: Object.freeze({
    name: 'All Words',
    redirect: '/german/all-words',
  }),
  GERMAN_RANDOM_GENERATOR: Object.freeze({
    name: 'Random Generator',
    redirect: '/german/random-generator',
  }),
  SEARCH_WORD: Object.freeze({
    name: 'Search Word',
    redirect: '/german/search-word',
  }),
  QUIZ: Object.freeze({
    name: 'Quiz',
    redirect: '/german/quiz',
  }),
  AUDIO_MOVER: Object.freeze({
    name: 'Audio Mover',
    redirect: '/german/audio-mover',
  }),
});

export const ROUTES = Object.freeze({
  HOMEPAGE: Object.freeze({
    name: 'Homepage',
    redirect: '/homepage',
  }),
  GERMAN_HOMEPAGE: Object.freeze({
    name: 'German Homepage',
    redirect: '/homepage/german',
  }),
  ...GERMAN_ROUTES, // Spread German routes here
});


export function getKeyFromName(name) {
  return Object.entries(ROUTES).find(
    ([key, route]) => route.name === name
  )?.[0]; // returns the key or undefined if not found
}

export const NO_BACK_BUTTON = ['/', ROUTES.HOMEPAGE]
