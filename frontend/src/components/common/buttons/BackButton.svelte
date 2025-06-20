<script>
    function goBack() {
        history.back();
    }
</script>

<button
        on:click={goBack}
        class="flex items-center text-blue-600 hover:text-blue-800 focus:outline-none"
        aria-label="Go back"
>
    <!-- Left arrow icon (SVG) -->
    <svg
            class="w-5 h-5 mr-2"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
    >
        <path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"></path>
    </svg>
    Back
</button>

