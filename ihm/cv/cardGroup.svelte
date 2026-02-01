<script>
    import { setContext } from 'svelte';
    import { writable } from 'svelte/store';

    const flippedCardId = writable(null);

    const handleFlip = (id) => {
        flippedCardId.update(currentId => {
            if (currentId === id) {
                return null;
            } else {
                return id;
            }
        });
    }

    setContext('card-group', {
        flippedCardId,
        handleFlip
    });
</script>

<slot />