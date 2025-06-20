<script>
  import { onMount } from 'svelte';
  //   import {  } from '../../../../wailsjs/go/main/MyApp'; // adjust import path
  import { ListFiles } from '../../../wailsjs/go/file_services/FileService'; // adjust import path

  let files = [];

  onMount(async () => {
    try {
      files = await ListFiles('storage/german/all-word');
    } catch (err) {
      alert('Failed to load files: ' + (err.message || err)); // Optional
    }
  });
</script>

<div
  class="flex flex-col items-center w-full max-w-md bg-white rounded-lg shadow-lg p-6 text-black"
>
  <div class="w-full max-w-xs mx-auto mt-10">
    <label
      for="translation-direction"
      class="block mb-2 text-sm font-semibold text-gray-900"
      >Choose direction:</label
    >
    <select
      id="translation-direction"
      class="block w-full p-3 border font-normal border-gray-500 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
    >
      <option value="de-en">German to English</option>
      <option value="en-de">English to German</option>
    </select>
  </div>

  <label
    for="search"
    class="block text-lg font-semibold mt-10 mb-2 text-gray-900">Search</label
  >
  <div class="w-full">
    <input
      id="search"
      type="text"
      placeholder="Type something..."
      class="w-full px-4 py-2 rounded border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
    />

    <ul
      id="autocomplete-list"
      class="max-w-md border border-gray-300 rounded mt-1 bg-white shadow-lg"
    >
      {#each files as file}
        <li class="p-2 hover:bg-gray-100 cursor-pointer">{file}</li>
      {/each}
    </ul>
  </div>
</div>
