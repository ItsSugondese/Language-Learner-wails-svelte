<script>
    import {onDestroy, onMount, tick} from 'svelte';
    import {ListDataFromFiles, ListFiles} from '../../../wailsjs/go/file_services/FileService';
    import {FilePathEnums} from '../../enums/file_path_enums.js';
    import {DelimiterConstant} from "../../constants/delimiter_constants.js";
    import {getSplitWordMeaningObject} from "../../helper/word/word_splitter.js";
    import {TranslateFromToEnums} from "../../enums/translate_from_to_enums.js";
    import {getRandomForList} from "../../utils/math_utils.js"; // adjust import path

    import { push } from 'svelte-spa-router';
    import {loadPageState, savePageState} from "../../utils/page_state.js";
    import {LogInfo} from "../../../wailsjs/runtime/runtime.js";
    import { SelectFile } from '../../../wailsjs/go/file_services/FilePicker.js';
    import InsertValueModal from "../../components/common/modals/InsertValueModal.svelte";
    let selectedFile = '';
    // on app load, redirect to last route if available
    const last = localStorage.getItem('lastRoute');
    if (last && last !== window.location.pathname) {
        push(last);
    }

    let searchText = '';
    let container; // Reference to wrapper div
    let allWords = []
    let direction = TranslateFromToEnums.GermanToEnglish
    let randomNum = 0


    async function chooseFile() {
        try {
            const filePath = await SelectFile();
            if (filePath) {
                allWords = await ListDataFromFiles(FilePathEnums.GermanAllWordFolderPath, filePath);
            }
        } catch (err) {
            console.error("File selection failed:", err);
        }
    }

    onMount(async () => {
        const state = await loadPageState();
        LogInfo("Loaded state: " + JSON.stringify(state));
        if (state) {
            searchText = state.searchText || '';
            allWords = state.allWords || [];
            direction = state.direction || direction;
            randomNum = state.randomNum || 0
            setWordMeaning(randomNum)
            // await tick();
        }


    });

    function setWordMeaning(num = getRandomForList(allWords)) {
        randomNum = num
        word = allWords[randomNum]
        const splittedWord = getSplitWordMeaningObject(word, DelimiterConstant.pipeSeparator, direction)
        if (splittedWord != null) {

            word = splittedWord.word;
            meaning = splittedWord.meaning;
        }
    }


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

    onDestroy(() => {
        savePageState({ direction, searchText, allWords, randomNum });
    });

    let showModal = false;
    let editorInput = 'Start here...';
    let finalText = '';

    function openModal() {
        showModal = true;
    }

    function handleDone(event) {
        finalText = event.detail.text;
        showModal = false;
    }

    function handleCancel() {
        showModal = false;
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
            Search
        </label>

        <div class="relative min-w-80">
            <div class="flex">

            <button on:click={() => showModal = true} class="p-2 bg-blue-600 text-white rounded">Insert Values</button>
            <button on:click={chooseFile} class="p-2 bg-blue-500 text-white rounded">Select File</button>
            </div>


            <InsertValueModal
                    bind:visible={showModal}
                    initialText={editorInput}
                    on:done={handleDone}
                    on:cancel={handleCancel}
            />
            {#if selectedFile}
                <p class="mt-2">Selected file: {selectedFile}</p>
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
