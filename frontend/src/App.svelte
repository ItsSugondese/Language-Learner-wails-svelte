<script>
  import Router, { location } from 'svelte-spa-router';

  import Homepage from './pages/homepages/Homepage.svelte';
  import GermanHomepage from './pages/homepages/GermanHomepage.svelte';
  import AllWords from './pages/german/AllWords.svelte';
  import RandomGenerator from './pages/german/RandomGenerator.svelte';
  import SearchWord from './pages/german/SearchWord.svelte';
  import Quiz from './pages/german/Quiz.svelte';
  import AudioMover from './pages/german/AudioMover.svelte';
  import { ROUTES, NO_BACK_BUTTON } from './constants/routes';
  import BackButton from './components/common/buttons/BackButton.svelte';

  const routes = {
    '/': Homepage, // Starting page
    [ROUTES.HOMEPAGE.redirect]: Homepage,
    [ROUTES.GERMAN_HOMEPAGE.redirect]: GermanHomepage,
    [ROUTES.GERMAN_ALL_WORDS.redirect]: AllWords,
    [ROUTES.GERMAN_RANDOM_GENERATOR.redirect]: RandomGenerator,
    [ROUTES.SEARCH_WORD.redirect]: SearchWord,
    [ROUTES.QUIZ.redirect]: Quiz,
    [ROUTES.AUDIO_MOVER.redirect]: AudioMover,
  };

  // Check if current route should show back button
  $: showBackButton = !NO_BACK_BUTTON.includes($location);

  $: localStorage.setItem('lastRoute', $location);
</script>

<div class="flex flex-col min-h-screen">
  <!-- Conditionally show BackButton -->
  {#if showBackButton}
    <div class="flex items-center p-4">
      <BackButton />
    </div>
  {/if}
  <div class="flex flex-1 flex-col items-center justify-center p-6">
    <Router {routes} restoreScrollState={false} />
  </div>
</div>
