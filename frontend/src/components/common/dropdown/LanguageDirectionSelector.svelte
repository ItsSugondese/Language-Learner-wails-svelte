<script>
    export let direction;
    export let onChange;
    export let label = 'Choose direction:';
    export let options = [
        { value: 'GermanToEnglish', label: 'German to English' },
        { value: 'EnglishToGerman', label: 'English to German' },
    ];
</script>

<div class="w-full max-w-xs mx-auto mt-10">
    <label
            for="translation-direction"
            class="block mb-2 text-sm font-semibold text-gray-900"
    >{label}</label
    >
    <select
            bind:value={direction}
            id="translation-direction"
            class="block w-full p-3 border font-normal border-gray-500 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            on:change={onChange}
    >
        {#each options as opt}
            <option value={opt.value}>{opt.label}</option>
        {/each}
    </select>
</div>
