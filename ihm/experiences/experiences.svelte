<script>
    import { onMount } from 'svelte';
    import Card from './card.svelte';

    let experiences = [];
    let error = '';

    async function getExperiences() {
        try {
            const response = await fetch('/experiences');
            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }
            const data = await response.json();
            experiences = data;
        } catch (e) {
            console.error("Échec de la récupération des expériences :", e.message || e);
            error = "Impossible de charger les expériences. Veuillez réessayer plus tard.";
        }
    }

    onMount(getExperiences);
</script>

<section class="experiences">
    <h2>Expériences professionnelles</h2>
    {#if error}
        <p class="error">{error}</p>
    {:else if experiences.length > 0}
        {#each experiences as experience, index}
            <Card {experience} {index} />
        {/each}
    {:else}
        <p>Chargement des expériences...</p>
    {/if}
</section>

<style>
    .experiences {
        margin: 2rem;
    }
    h2 {
        font-size: 2.5rem;
        font-weight: 700;
        text-align: center;
        padding-bottom: 0.5rem;
        margin-bottom: 2rem;
    }
    .error {
        color: red;
        text-align: center;
    }
</style>
