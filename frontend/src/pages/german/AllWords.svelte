<script>
  import {onDestroy, onMount, tick} from 'svelte';

  import {
    ListDataFromFilePathEnum,
    ListFiles,
    LoadAllAudioInDirectory,
  } from '../../../wailsjs/go/file_services/FileService';
  import {FilePathEnums} from '../../enums/file_path_enums.js';
  import {TranslateFromToEnums} from '../../enums/translate_from_to_enums.js';
  import {location} from 'svelte-spa-router';
  import {loadPageState, savePageState} from '../../utils/page_state.js';
  import {LogInfo} from '../../../wailsjs/runtime/runtime.js';
  import WordMeaningDisplayCard from "../../components/common/card/WordMeaningDisplayCard.svelte";
  import LanguageDirectionSelector from "../../components/common/dropdown/LanguageDirectionSelector.svelte";
  import SuccessPopup from "../../components/common/popup/SuccessPopup.svelte";

  let searchText = '';
  let files = [];
  let isOpen = false;
  let highlightedIndex = -1;
  let container; // Reference to wrapper div
  let allWords = [];
  let direction;
  let randomNum = 0;
  let currentLocation;
  let wordMeaningCardRef;
  let showSuccessPopup = false;


  onMount(async () => {
    direction = TranslateFromToEnums.GermanToEnglish.name;
    currentLocation = $location
    const state = await loadPageState(currentLocation);
    if (state) {
      searchText = state.searchText || '';
      allWords = state.allWords || [];
      direction = state.direction || direction;
      randomNum = state.randomNum || 0;

      await tick();
      if (wordMeaningCardRef) {
        wordMeaningCardRef.setWordMeaning(randomNum);
      } else {
        LogInfo("wordMeaningCardRef not available yet");
      }
    }

    try {
      files = await ListFiles(FilePathEnums.GermanAllWordFolderPath, false);
    } catch (err) {
      alert('Failed to load files: ' + (err.message || err));
    }
    window.addEventListener('click', handleClickOutside);

  });

  function handleRandomNumChange(event) {
    randomNum = event.detail.randomNum;
  }

  function handleWordRemoved(event) {
    allWords = event.detail.allWords;
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
    allWords = await ListDataFromFilePathEnum(
      FilePathEnums.GermanAllWordFolderPath,
      file,
    );
    wordMeaningCardRef.setWordMeaning();

    isOpen = false; // close dropdown
  }

  async function loadAllWordsAudio() {
    await LoadAllAudioInDirectory(FilePathEnums.GermanAllWordFolderPath,FilePathEnums.GermanAllWordAudioFolderAbsolutePath, searchText);
    showSuccessPopup = true;
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



  const directionOptions = [
    { value: TranslateFromToEnums.GermanToEnglish.name, label: 'German to English' },
    { value: TranslateFromToEnums.EnglishToGerman.name, label: 'English to German' },
  ];

  function handleDirectionChange(event) {
    wordMeaningCardRef.swapWordAndMeaning();
  }

  onDestroy(() => {
    window.removeEventListener('click', handleClickOutside);

    savePageState({ direction, searchText, allWords, randomNum }, currentLocation);
  });


  // Reactive filtered list
  $: filteredFiles = files.filter((file) =>
          file.toLowerCase().includes(searchText.toLowerCase()),
  );

  // Auto-highlight first match when filtered list changes
  $: if (filteredFiles.length > 0) {
    highlightedIndex = 0;
  }

</script>

<div
  class="flex flex-1 flex-col items-center bg-white rounded-lg shadow-lg p-6 text-black space-y-14"
>
  <div class="top">
    <LanguageDirectionSelector
            bind:direction
            onChange={handleDirectionChange}
            options={directionOptions}
    />

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

    <button
            on:click={loadAllWordsAudio}
            class="bg-blue-500 hover:bg-blue-600 text-white font-semibold px-4 py-2 rounded mt-3"
    >
      Load Audio
    </button>
  </div>

  <WordMeaningDisplayCard
          bind:this={wordMeaningCardRef}
          allWords={allWords}
          direction="{direction}"
          on:randomNumChanged={handleRandomNumChange}
          on:wordRemoved={handleWordRemoved}
          topic="{searchText}"
  />
</div>

{#if showSuccessPopup}
  <SuccessPopup message="Your action was successful!" onClose={() => showSuccessPopup = false} />
{/if}
