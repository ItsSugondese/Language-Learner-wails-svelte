<script>
  import {onDestroy, onMount, tick} from 'svelte';
  import { ListDataFromFilePathEnum, ListFiles} from '../../../wailsjs/go/file_services/FileService';
  import {FilePathEnums} from '../../enums/file_path_enums.js';
  import {DelimiterConstant} from "../../constants/delimiter_constants.js";
  import {getSplitWordMeaningObject} from "../../helper/word/word_splitter.js";
  import {TranslateFromToEnums} from "../../enums/translate_from_to_enums.js";
  import {getRandomForList} from "../../utils/math_utils.js"; // adjust import path

  import {location} from 'svelte-spa-router';
  import {loadPageState, savePageState} from "../../utils/page_state.js";
  import {LogInfo} from "../../../wailsjs/runtime/runtime.js";


  let searchText = '';
  let files = [];
  let isOpen = false;
  let highlightedIndex = -1;
  let container; // Reference to wrapper div
  let allWords = []
  let direction = TranslateFromToEnums.GermanToEnglish
  let randomNum = 0

  onMount(async () => {
    const state = await loadPageState($location);
    LogInfo("Loaded state: " + JSON.stringify(state));
    if (state) {
      searchText = state.searchText || '';
      allWords = state.allWords || [];
      direction = state.direction || direction;
      randomNum = state.randomNum || 0
      setWordMeaning(randomNum)
      // await tick();
    }
    window.addEventListener('click', handleClickOutside);

    try {
      files = await ListFiles(FilePathEnums.GermanAllWordFolderPath, false);
    } catch (err) {
      alert('Failed to load files: ' + (err.message || err));
    }
  });

  // Reactive filtered list
  $: filteredFiles = files.filter((file) =>
    file.toLowerCase().includes(searchText.toLowerCase()),
  );

  // Auto-highlight first match when filtered list changes
  $: if (filteredFiles.length > 0) {
    highlightedIndex = 0;
  }
  const toggleDropdown = () => {
    isOpen = true;
  };

  function handleClickOutside(event) {
    if (!container.contains(event.target)) {
      isOpen = false;
    }
  }

  // When user clicks a file from the list:
  async function selectFile(file) {
    searchText = file;
    allWords = await ListDataFromFilePathEnum(FilePathEnums.GermanAllWordFolderPath, file);
    setWordMeaning()

    isOpen = false; // close dropdown
  }

  function setWordMeaning(num = getRandomForList(allWords)) {
    randomNum = num
    word = allWords[randomNum]
    const splittedWord = getSplitWordMeaningObject(word, DelimiterConstant.pipeSeparator, direction)
    if (splittedWord != null) {

      word = splittedWord.word;
      meaning = splittedWord.meaning;
    }
  }

  function handleKeyDown(e) {
    if (e.key === 'ArrowDown') {
      e.preventDefault();
      highlightedIndex = (highlightedIndex + 1) % filteredFiles.length;
    } else if (e.key === 'ArrowUp') {
      e.preventDefault();
      highlightedIndex =
        (highlightedIndex - 1 + filteredFiles.length) % filteredFiles.length;
    } else if (e.key === 'Enter' && filteredFiles.length > 0 && isOpen) {
      e.preventDefault();
      selectFile(filteredFiles[highlightedIndex]);
    }
  }

  onDestroy(() => {
    window.removeEventListener('click', handleClickOutside);

    savePageState({ direction, searchText, allWords, randomNum }, $location);
  });

  let word = null;
  let meaning = null;
  let showMeaning = true;

  function reveal() {
    showMeaning = !showMeaning;
  }

  function nextCard() {
    // Replace this with your actual word list logic
    setWordMeaning()
  }

  function handleDirectionChange(event) {
    const temp = word;
    word = meaning;
    meaning = temp;
  }


</script>

<div
  class="flex flex-1 flex-col items-center bg-white rounded-lg shadow-lg p-6 text-black space-y-14"
>
  <div class="top">
    <div class="w-full max-w-xs mx-auto mt-10">
      <label
        for="translation-direction"
        class="block mb-2 text-sm font-semibold text-gray-900"
        >Choose direction:</label
      >
      <select
              bind:value={direction}
        id="translation-direction"
        class="block w-full p-3 border font-normal border-gray-500 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        on:change={handleDirectionChange}>
        <option value={TranslateFromToEnums.GermanToEnglish}>German to English</option>
        <option value={TranslateFromToEnums.EnglishToGerman}>English to German</option>
      </select>
    </div>

    <label
      for="search"
      class="block text-lg font-semibold mt-10 mb-2 text-gray-900"
    >
      {allWords.length} Words
    </label>

    <div class="relative min-w-80">
      <input
        id="search"
        type="text"
        bind:value={searchText}
        placeholder="Type something..."
        class="w-full px-4 py-2 rounded border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
        on:input={() => (isOpen = true)}
        on:keydown={handleKeyDown}
        on:click={toggleDropdown}
        bind:this={container}
      />

      {#if isOpen}
        {#if filteredFiles.length > 0}
          <ul
            id="autocomplete-list"
            class="absolute z-10 w-full max-h-48 overflow-y-auto border border-gray-300 rounded mt-1 bg-white shadow-lg"
          >
            {#each filteredFiles as file, i}
              <li
                class={`p-2 cursor-pointer rounded transition-colors duration-150
            ${i === highlightedIndex ? 'bg-blue-100 text-blue-800' : 'hover:bg-gray-100'}`}
                on:click={() => selectFile(file)}
                on:keydown={(e) => {
                  if (e.key === 'Enter' || e.key === ' ') {
                    e.preventDefault();
                    selectFile(file);
                  }
                }}
              >
                {file}
              </li>
            {/each}
          </ul>
        {:else if files.length > 0}
          <p class="text-sm text-gray-500 mt-2">No results found.</p>
        {/if}
      {/if}
    </div>
  </div>

  {#if word !== null}
  <div class="bottom flex items-center justify-center bg-gray-100">
    <div
      class="bg-white rounded-xl shadow-lg p-6 w-full max-w-md text-center space-y-6"
    >
      <div class="space-y-16">
        <div class="text-4xl font-bold text-blue-600">{word}</div>

        <div class="space-x-2">
          <button
            on:click={reveal}
            class="bg-blue-500 hover:bg-blue-600 text-white font-semibold px-4 py-2 rounded"
          >
            {#if showMeaning}
              Hide Meaning
            {:else}
              Show Meaning
            {/if}
          </button>

          <button
            on:click={nextCard}
            class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-semibold px-4 py-2 rounded"
          >
            Next Word
          </button>
        </div>
      </div>

      {#if showMeaning}
        <div class="text-xl text-green-600">{meaning}</div>
      {/if}
    </div>
  </div>
    {/if}
</div>
