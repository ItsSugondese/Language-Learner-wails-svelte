<script>
  import { push } from 'svelte-spa-router';
  import { ROUTES } from '../../constants/routes.js';
  import NavigationButton from '../../components/common/buttons/NavigationButton.svelte';
  // You can replace or extend this list
  let languages = ['German'];

  function selectLanguage(lang) {
    // push(ROUTES.GERMAN_HOMEPAGE);
    push(ROUTES.GERMAN_HOMEPAGE.redirect);
  }
</script>

<div
  class="flex flex-col items-center w-full max-w-md bg-white rounded-lg shadow-lg p-6"
>
  <h1 class="text-2xl font-bold text-gray-800 mb-6">Select language</h1>

  <div class="flex flex-col gap-4 w-full max-w-xs">
    {#each languages as lang}
      <NavigationButton on:click={() => selectLanguage(lang)}
        >{lang}</NavigationButton
      >
    {/each}
  </div>
</div>
