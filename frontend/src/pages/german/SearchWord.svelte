<script>
    import {onDestroy, onMount} from 'svelte';

    import {ListDataFromAllFiles} from '../../../wailsjs/go/file_services/FileService';
    import {FilePathEnums} from '../../enums/file_path_enums.js';
    import {TranslateFromToEnums} from '../../enums/translate_from_to_enums.js';
    import LanguageDirectionSelector from "../../components/common/dropdown/LanguageDirectionSelector.svelte";
    import {DelimiterConstant} from "../../constants/delimiter_constants.js";
    import {normalizeUmlauts} from "../../utils/string_utils.js";

    let searchText = '';
    let isOpen = false;
    let highlightedIndex = -1;
    let container; // Reference to wrapper div
    let allWords = [];
    let words = [];
    let direction = TranslateFromToEnums.GermanToEnglish.name;
    let meaning =null

    const genericValueGermanMap = {};
    const genericValueEnglishMap = {};


    onMount(async () => {

        try {
            allWords = await ListDataFromAllFiles(FilePathEnums.GermanAllWordFolderPath);
            allWords.forEach(e => {
                const splittedWord = e.split(DelimiterConstant.pipeSeparator);
                if (splittedWord.length >= 2) {
                    let german = splittedWord[0].trim();
                    let english = splittedWord[1].trim();

                    let tempGerman = german;
                    let tempEnglish = english;
                    if (german in genericValueGermanMap) {
                        tempEnglish = english + " ; " + genericValueGermanMap[german];
                    }

                    if (english in genericValueEnglishMap) {
                        tempGerman = german + " ; " + genericValueEnglishMap[english];
                    }
                    genericValueGermanMap[german] = tempEnglish;
                    genericValueEnglishMap[english] = tempGerman;
                }
            });

            words = Object.keys(genericValueGermanMap);
        } catch (err) {
            alert('Failed to load files: ' + (err.message || err));
        }
        window.addEventListener('click', handleClickOutside);

    });



    const toggleDropdown = () => {
        isOpen = true;
    };

    function handleClickOutside(event) {
        if (!container.contains(event.target)) {
            isOpen = false;
        }
    }

    // When user clicks a file from the list:
    async function selectedWord(word) {
        searchText = word;
        if (direction === TranslateFromToEnums.GermanToEnglish.name) {
            meaning = genericValueGermanMap[word];
        } else {
            meaning = genericValueEnglishMap[word];
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
            selectedWord(filteredFiles[highlightedIndex]);
            isOpen = false;
        }
    }



    const directionOptions = [
        { value: TranslateFromToEnums.GermanToEnglish.name, label: 'German to English' },
        { value: TranslateFromToEnums.EnglishToGerman.name, label: 'English to German' },
    ];

    function handleDirectionChange(event) {
        if (direction === TranslateFromToEnums.GermanToEnglish.name) {
            words = Object.keys(genericValueGermanMap);
        } else {
            words = Object.keys(genericValueEnglishMap);
        }

        searchText = ''
        meaning = null;

    }

    onDestroy(() => {
        window.removeEventListener('click', handleClickOutside);
    });


    // Reactive filtered list
    $: filteredFiles = words.filter((word) =>
        normalizeUmlauts(word).includes(normalizeUmlauts(searchText))
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
            Search
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
                                    on:click={() => selectedWord(file)}
                                    on:keydown={(e) => {
                  if (e.key === 'Enter' || e.key === ' ') {
                    e.preventDefault();
                    selectedWord(file);
                  }
                }}
                            >
                                {file}
                            </li>
                        {/each}
                    </ul>
                {:else if words.length > 0}
                    <p class="text-sm text-gray-500 mt-2">No results found.</p>
                {/if}
            {/if}
        </div>

    </div>




    {#if meaning !== null}
    <div class="bottom flex items-center justify-center bg-gray-100">
        <div
                class="bg-white rounded-xl shadow-lg p-6 w-full max-w-md text-center space-y-6"
        >
                <div
                        class="text-4xl font-bold text-blue-600 flex items-center justify-center space-x-2"
                >
                    <span>{meaning}</span>
                </div>




        </div>
    </div>
    {/if}

</div>


