<script>
  import {createEventDispatcher, onDestroy, onMount} from 'svelte';
  import { LogInfo } from '../../../../wailsjs/runtime/runtime.js';
  import { FindFileFromFilePathEnumIfExistsAndReturnEncoded } from '../../../../wailsjs/go/file_services/FileService.js';
  import { FilePathEnums } from '../../../enums/file_path_enums.js';
  import { TranslateFromToEnums } from '../../../enums/translate_from_to_enums.js';
  import {getRandomForList} from "../../../utils/math_utils.js";
  import {getSplitWordMeaningObject} from "../../../helper/word/word_splitter.js";
  import {DelimiterConstant} from "../../../constants/delimiter_constants.js";
  const dispatch = createEventDispatcher();
  export let word = '';
  export let meaning = '';
  export let direction;
  export let allWords = [];
  let showMeaning = true;
  export let randomNum = 0;

  export let isRandomScreen = false;
  export let topic = '';

  let removeButton;
  let nextButton;
  let wordAudio;
  let meaningAudio;


  onMount( () => {
    window.addEventListener('keydown', handleKeyDown);

  })

  function reveal() {
    showMeaning = !showMeaning;
  }

  function nextCard() {
    setWordMeaning();
  }

  export function setWordMeaning(num = getRandomForList(allWords)) {
    if(allWords.length !== 0) {
      randomNum = num;
      word = allWords[randomNum];
      const splittedWord = getSplitWordMeaningObject(
              word,
              DelimiterConstant.pipeSeparator,
              direction,
      );
      if (splittedWord != null) {
        word = splittedWord.word;
        meaning = splittedWord.meaning;
      }

      dispatch('randomNumChanged', {randomNum});
    }
  }

  export function getWordMeaningValue() {
    return Object.freeze({
      word: word,
      meaning: meaning,
    });
  }

  export function swapWordAndMeaning() {
    const temp = word;
    word = meaning;
    meaning = temp;
  }
  function removeWord() {
    allWords.splice(randomNum, 1);

    dispatch('wordRemoved', { allWords });
    setWordMeaning();
  }

  function handleKeyDown(event) {
    if (event.altKey && event.key === 'r') {
      event.preventDefault(); // prevent default browser save
      removeWord()
    }
    if (event.key === 'ArrowRight') {
      event.preventDefault(); // prevent default browser save
      nextCard()
    }

    if (event.key === 'ArrowUp') {
      event.preventDefault(); // prevent default browser save
      playAudio(word, true)
    }

    if (event.key === 'ArrowDown') {
      event.preventDefault(); // prevent default browser save
      playAudio(meaning, false)
    }
  }


  async function playAudio(word, isWord) {
    if (word) {
      let base64;

      if (direction === TranslateFromToEnums.EnglishToGerman.name) {
        if (isWord) {
          base64 = await FindFileFromFilePathEnumIfExistsAndReturnEncoded(
                  FilePathEnums.EnglishAudioFolderAbsolutePath,
                  '',
                  word,
                  false,
          );
        } else {
          base64 = await FindFileFromFilePathEnumIfExistsAndReturnEncoded(
                  FilePathEnums.GermanAudioFolderAbsolutePath,
                  !isRandomScreen ? topic : '',
                  word,
                  isRandomScreen,
          );
        }
      } else {
        if (isWord) {
          base64 = await FindFileFromFilePathEnumIfExistsAndReturnEncoded(
                  FilePathEnums.GermanAudioFolderAbsolutePath,
                  !isRandomScreen ? topic : '',
                  word,
                  isRandomScreen,
          );
        } else {
          base64 = await FindFileFromFilePathEnumIfExistsAndReturnEncoded(
                  FilePathEnums.EnglishAudioFolderAbsolutePath,
                  '',
                  word,
                  false,
          );
        }


      }
      if (base64) {
        const audio = new Audio(base64);
        audio.play();
      } else {
        LogInfo("Audio file not found.");
      }
    }
  }

  onDestroy(() => {
    window.removeEventListener('keydown', handleKeyDown);
  });
</script>

{#if word !== null}
  <div class="bottom flex items-center justify-center bg-gray-100">
    <div
            class="bg-white rounded-xl shadow-lg p-6 w-full max-w-md text-center space-y-6"
    >
      <div class="space-y-16">
        <div
                class="text-4xl font-bold text-blue-600 flex items-center justify-center space-x-2"
        >
          <span>{word}</span>
          <button
                  on:click={() => playAudio(word, true)}
                  class="text-blue-600 hover:text-blue-800"
          >
            🔊
          </button>
        </div>

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
                  on:click={removeWord}
                  class="bg-red-500 hover:bg-gray-400 text-white font-semibold px-4 py-2 rounded"
          >
            Remove
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
        <div
                class="text-xl text-green-600 flex items-center justify-center space-x-2"
        >
          <span>{meaning}</span>
          <button
                  on:click={() => playAudio(meaning, false)}
                  class="text-green-600 hover:text-green-800"
          >
            🔊
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}
