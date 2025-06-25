<script>
    export let message = "Success!";
    export let onClose = () => {};

    function handleClose() {
        onClose(); // Notify parent
    }


</script>

<div class="fixed top-5 right-5 bg-green-500 text-white px-5 py-3 rounded-xl shadow-lg flex items-center space-x-4 z-50">
    <span>{message}</span>
    <button
            on:click={handleClose}
            class="bg-white text-green-700 font-semibold px-3 py-1 rounded hover:bg-gray-100 transition"
    >
        Okay
    </button>
</div>
