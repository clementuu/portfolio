<script>
    import { onMount } from 'svelte';
    import Card from './card.svelte';

    let formations = [];
    let error = '';

    async function getFormations() {
        try {
            const response = await fetch('/formations');
            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }
            const data = await response.json();
            formations = data;
        } catch (e) {
            console.error("Échec de la récupération des formations :", e.message || e);
            error = "Impossible de charger les formations. Veuillez réessayer plus tard.";
        }
    }

    onMount(getFormations);
</script>

<section class="formations">
    <h2>Formations</h2>
    {#if error}
        <p class="error">{error}</p>
    {:else if formations.length > 0}
        {#each formations as formation, index}
            <Card {formation} {index} />
        {/each}
    {:else}
        <p>Chargement des formations...</p>
    {/if}
</section>

<style>
    .formations {
        margin: 2rem;
    }
    h2 {
        font-size: 2.5rem; /* Larger font size */
        font-weight: 700; /* Bolder */
        text-align: center; /* Center the title */
        padding-bottom: 0.5rem;
        margin-bottom: 2rem; /* More space below the title */
    }
    .error {
        color: red;
        text-align: center;
    }
</style>
