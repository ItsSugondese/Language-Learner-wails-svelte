<script>
  import {onDestroy, onMount, tick} from 'svelte';
  import {ListDataFromFilePath} from '../../../wailsjs/go/file_services/FileService';
  import {DelimiterConstant} from '../../constants/delimiter_constants.js';
  import {TranslateFromToEnums} from '../../enums/translate_from_to_enums.js';
  import {loadPageState, savePageState} from '../../utils/page_state.js';
  import {LogInfo} from '../../../wailsjs/runtime/runtime.js';
  import {SelectFile} from '../../../wailsjs/go/file_services/FilePicker.js';
  import InsertValueModal from '../../components/common/modals/InsertValueModal.svelte';
  import {convertListToString, convertStringToList,} from '../../utils/string_utils.js';
  import {location} from 'svelte-spa-router';
  import WordMeaningDisplayCard from "../../components/common/card/WordMeaningDisplayCard.svelte";
  import LanguageDirectionSelector from "../../components/common/dropdown/LanguageDirectionSelector.svelte";


  let allWords = [];
  let direction = TranslateFromToEnums.GermanToEnglish.name;
  let randomNum = 0;
  let currentLocation

  let wordMeaningCardRef;

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
  });


  function handleRandomNumChange(event) {
    randomNum = event.detail.randomNum;
  }

  function handleWordRemoved(event) {
    allWords = event.detail.allWords;
  }



  const directionOptions = [
    { value: TranslateFromToEnums.GermanToEnglish.name, label: 'German to English' },
    { value: TranslateFromToEnums.EnglishToGerman.name, label: 'English to German' },
  ];

  function handleDirectionChange(event) {
    wordMeaningCardRef.swapWordAndMeaning();
  }


  let showModal = false;
  let editorInput = '';

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

  onDestroy(() => {
    savePageState({ direction, allWords, randomNum }, currentLocation);
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

    <WordMeaningDisplayCard
            bind:this={wordMeaningCardRef}
            allWords={allWords}
            direction="{direction}"
            on:randomNumChanged={handleRandomNumChange}
            on:wordRemoved={handleWordRemoved}
            isRandomScreen="{true}"
    />



</div>
