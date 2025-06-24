<script>
  import {onDestroy, onMount, tick} from 'svelte';
  import { ListDataFromFilePath } from '../../../wailsjs/go/file_services/FileService';
  import { DelimiterConstant } from '../../constants/delimiter_constants.js';
  import { getSplitWordMeaningObject } from '../../helper/word/word_splitter.js';
  import { TranslateFromToEnums } from '../../enums/translate_from_to_enums.js';
  import { getRandomForList } from '../../utils/math_utils.js'; // adjust import path
  import { loadPageState, savePageState } from '../../utils/page_state.js';
  import { LogInfo } from '../../../wailsjs/runtime/runtime.js';
  import { SelectFile } from '../../../wailsjs/go/file_services/FilePicker.js';
  import InsertValueModal from '../../components/common/modals/InsertValueModal.svelte';
  import {
    convertListToString,
    convertStringToList,
  } from '../../utils/string_utils.js';
  import { location } from 'svelte-spa-router';
  import WordMeaningDisplayCard from "../../components/common/card/WordMeaningDisplayCard.svelte";
  import LanguageDirectionSelector from "../../components/common/dropdown/LanguageDirectionSelector.svelte";


  let searchText = '';
  let allWords = [];
  let direction = TranslateFromToEnums.GermanToEnglish.name;
  let randomNum = 0;
  let currentLocation

  async function chooseFile() {
    try {
      const filePath = await SelectFile();
      if (filePath) {
        const newWords = await ListDataFromFilePath(filePath);
        allWords = [...allWords, ...newWords];
      }
    } catch (err) {
      console.error('File selection failed:', err);
    }
  }

  onMount(async () => {
    currentLocation = $location

    const state = await loadPageState(currentLocation);
    if (state) {
      searchText = state.searchText || '';
      allWords = state.allWords || [];
      direction = state.direction || direction;
      randomNum = state.randomNum || 0;
      setWordMeaning(randomNum);
      // await tick();
    }
  });

  function setWordMeaning(num = getRandomForList(allWords)) {
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
  }

  let word = null;
  let meaning = null;


  function handleNext() {
    // Replace this with your actual word list logic
    setWordMeaning();
  }


  const directionOptions = [
    { value: TranslateFromToEnums.GermanToEnglish.name, label: 'German to English' },
    { value: TranslateFromToEnums.EnglishToGerman.name, label: 'English to German' },
  ];

  function handleDirectionChange(event) {
    const temp = word;
    word = meaning;
    meaning = temp;
  }


  let showModal = false;
  let editorInput = 'Start here...';

  function openModal() {
    editorInput = convertListToString(allWords, DelimiterConstant.lineBreak);
    LogInfo(editorInput);
    showModal = true;
  }

  function handleDone(event) {
    allWords = convertStringToList(
      event.detail.text,
      DelimiterConstant.lineBreak,
    );
    showModal = false;
  }

  function handleCancel() {
    showModal = false;
  }

  function speak(text) {
    const utterance = new SpeechSynthesisUtterance(text);
    utterance.lang = 'de-DE'; // change to 'en-US' for English
    speechSynthesis.speak(utterance);
  }
  onDestroy(() => {
    savePageState({ direction, searchText, allWords, randomNum }, currentLocation);
  });
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
      <div class="flex w-full justify-around">
        <button on:click={openModal} class="p-2 bg-blue-600 text-white rounded"
          >Insert Values</button
        >
        <button on:click={chooseFile} class="p-2 bg-blue-500 text-white rounded"
          >Select File</button
        >
      </div>

      <InsertValueModal
        bind:visible={showModal}
        initialText={editorInput}
        on:done={handleDone}
        on:cancel={handleCancel}
      />
    </div>
  </div>

  {#if word !== null}
    <WordMeaningDisplayCard
            word={word}
            meaning={meaning}
            on:next={handleNext}
            direction="{direction}"
    />
  {/if}
</div>
