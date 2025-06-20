<script>
  // Function to go back in browser history
  import { push } from 'svelte-spa-router';
  import {
    GERMAN_ROUTES,
    getKeyFromName,
    ROUTES,
  } from '../../constants/routes.js';
  import NavigationButton from '../../components/common/buttons/NavigationButton.svelte';
  // You can replace or extend this list
  let options = Object.values(GERMAN_ROUTES).map((route) => route.name);

  function selectLanguage(option) {
    push(ROUTES[getKeyFromName(option)].redirect);
  }
</script>

<div
  class="flex flex-col items-center w-full max-w-md bg-white rounded-lg shadow-lg p-6"
>
  <h1 class="text-2xl font-bold text-gray-800 mb-6">Select language</h1>

  <div class="flex flex-col gap-4 w-full max-w-xs">
    {#each options as option}
      <NavigationButton on:click={() => selectLanguage(option)}
        >{option}</NavigationButton
      >
    {/each}
  </div>
</div>
