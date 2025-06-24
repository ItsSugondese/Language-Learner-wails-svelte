<script>
  import { createEventDispatcher } from 'svelte';
  import { LogInfo } from '../../../../wailsjs/runtime/runtime.js';
  import { FindFileFromFilePathEnumIfExistsAndReturnEncoded } from '../../../../wailsjs/go/file_services/FileService.js';
  import { FilePathEnums } from '../../../enums/file_path_enums.js';
  import { TranslateFromToEnums } from '../../../enums/translate_from_to_enums.js';
  const dispatch = createEventDispatcher();
  export let word = '';
  export let meaning = '';
  export let direction;
  let showMeaning = true;

  function reveal() {
    showMeaning = !showMeaning;
  }

  function nextCard() {
    dispatch('next');
  }

  async function playAudio(word, isWord) {
    let base64;

    if (direction === TranslateFromToEnums.EnglishToGerman.name) {
      if (isWord) {
        base64 = await FindFileFromFilePathEnumIfExistsAndReturnEncoded(
          FilePathEnums.EnglishAudioFolderAbsolutePath,
          '',
          word + '.mp3',
          false,
        );
      } else {
        base64 = await FindFileFromFilePathEnumIfExistsAndReturnEncoded(
          FilePathEnums.GermanAudioFolderAbsolutePath,
          '',
          word + '.mp3',
          true,
        ); // e.g. /home/user/sound.mp3
      }
    } else {
      if (isWord) {
        base64 = await FindFileFromFilePathEnumIfExistsAndReturnEncoded(
          FilePathEnums.GermanAudioFolderAbsolutePath,
          '',
          word + '.mp3',
          true,
        ); // e.g. /home/user/sound.mp3
      } else {
        base64 = await FindFileFromFilePathEnumIfExistsAndReturnEncoded(
          FilePathEnums.EnglishAudioFolderAbsolutePath,
          '',
          word + '.mp3',
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
</script>

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
