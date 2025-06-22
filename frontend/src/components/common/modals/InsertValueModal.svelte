<script>
    import CodeMirror from "svelte-codemirror-editor";
    import { createEventDispatcher, onMount } from 'svelte';

    export let visible = false;
    export let initialText = '';

    let text = '';
    const dispatch = createEventDispatcher();

    onMount(() => {
        text = initialText;
    });

    const handleDone = () => dispatch('done', { text });
    const handleCancel = () => dispatch('cancel');
    const handleClear = () => text = '';

    // CodeMirror options
    const options = {
        lineNumbers: true,
        lineWrapping: true,  // enable wrapping so line numbers count wrapped lines properly
        mode: 'text', // plain text mode
        theme: 'default', // you can customize
    };
</script>

{#if visible}
    <div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
        <div class="bg-white dark:bg-gray-800 rounded-xl shadow-xl w-full max-w-3xl flex flex-col">
            <!-- Header -->
            <div class="px-6 py-4 border-b border-gray-200 dark:border-gray-700">
                <h2 class="text-xl font-semibold text-gray-800 dark:text-gray-100">Editor</h2>
            </div>

            <!-- Body -->
            <div class="p-4 flex-1 overflow-auto" style="height:400px;  text-align: left;">
                    <CodeMirror bind:value={text} options={{ lineNumbers: true, lineWrapping: true }} style="height: 100%; " />

            </div>

            <!-- Footer -->
            <div class="px-6 py-4 border-t border-gray-200 dark:border-gray-700 flex justify-end gap-3">
                <button on:click={handleDone} class="bg-blue-600 hover:bg-blue-700 text-white font-semibold px-4 py-2 rounded">OK</button>
                <button class="bg-yellow-500 hover:bg-yellow-600 text-white font-semibold px-4 py-2 rounded">Filter</button>
                <button on:click={handleClear} class="bg-red-500 hover:bg-red-600 text-white font-semibold px-4 py-2 rounded">Clear</button>
                <button on:click={handleCancel} class="bg-gray-400 hover:bg-gray-500 text-white font-semibold px-4 py-2 rounded">Cancel</button>
            </div>
        </div>
    </div>
{/if}


