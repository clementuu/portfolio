<script>
    export let className;
</script>

<article class = {className}>
    <div>
        <h3>
            <slot name="title" />
        </h3>
        <p>
            <strong>
                <slot name="subtitle" />
            </strong>
        </p>
    </div>
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

    h3 {
        font-size: 1.3rem;
        color: #333;
        margin-top: 0;
        margin-bottom: 0.5rem;
        border-bottom: none;
    }

    p {
        font-size: 1rem;
        line-height: 1.5;
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
