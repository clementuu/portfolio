<script>
    import { getContext } from 'svelte';

    export let id;

    const { flippedCardId, handleFlip } = getContext('card-group');

    $: isFlipped = $flippedCardId === id;
</script>

<article class:flipped={isFlipped}>
    <button
        class="invisible-button"
        on:click={() => {
            handleFlip(id);
        }}
    >
        {#if isFlipped}
            <slot name="upside-down"/>
        {:else}
            <slot name="hawkins"/>
        {/if}
    </button>
</article>

<style>
    article {
        background-color: white;
        padding: 1.5rem;
        margin-bottom: 1.5rem;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        transition: background-color 0.3s ease-in-out, color 0.3s ease-in-out, box-shadow 0.2s ease-in-out, transform 0.2s ease-in-out;
        cursor: pointer; /* Indicates the element is clickable */
    }

    article:hover {
        transform: scale(1.01); /* Provides visual feedback on hover */
        box-shadow: 0 6px 10px rgba(0, 0, 0, 0.15);
    }

    /* Style for the "upside-down" state, applied on click */
    article.flipped {
        background-color: #363636;
        color: white;
    }

    .invisible-button {
        all: unset;
        display: block;
        width: 100%;
        height: 100%;
        cursor: pointer;
    }


    /* --- Simplified global styles for slotted content --- */
    /* Default text color for "hawkins" content */
    :global(article ul),
    :global(article li),
    :global(article p) {
        transition: color 0.3s ease-in-out;
        color: #666;
    }

    
    :global(article h3),
    :global(article h4) {
        transition: color 0.3s ease-in-out;
        color: black;
    }

    /* Text color for "upside-down" content */
    article.flipped :global(ul),
    article.flipped :global(li),
    article.flipped :global(h3),
    article.flipped :global(h4),
    article.flipped :global(p) {
        color: white;
    }
</style>
