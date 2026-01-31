<script>
    export let className;
</script>

<article class = {className}>
    <slot />
</article>

<style>
    article {
        background-color: #f9f9f9;
        padding: 1.5rem;
        margin-bottom: 1.5rem;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        transition: box-shadow 0.2s ease-in-out;
    }

    .flex {
        display: flex;
        justify-content: space-between;
    }

    /* Styles for slotted content */
    :global(article ul) {
        padding-left: 20px;
        margin-top: 1rem;
        color: #666;
    }

    :global(article li) {
        margin-bottom: 0.5rem;
    }
</style>
