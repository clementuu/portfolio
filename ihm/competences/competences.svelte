<svelte:options customElement="competences-portfolio" />

<script>
    import Card from './card.svelte';
    import { onMount } from 'svelte';

    let error = null;
    let allCompetences = [];
    let displayedCompetences = null; // compétences à afficher après filtrage

    async function getCompetences() {
        try {
            const response = await fetch('/competences');

            if (!response.ok) {
                error = response.statusText
                throw new Error(`Erreur HTTP ${response.status}: ${error}`);
            }

            const data = await response.json();
            allCompetences = data;
        } catch (e) {
            error = e;
            console.error("Échec de la récupération des compétences :", error.message || error);
        }
    }

    function filterCompetences() {
        const urlParams = new URLSearchParams(window.location.search);
        const type = urlParams.get('type');

        if (type) {
            displayedCompetences = allCompetences.filter(comp => comp.type.toLowerCase() === type);
        } else {
            displayedCompetences = allCompetences;
        }
    }
    
    onMount(async () => {
        await getCompetences();
        filterCompetences(); // filtre les données après le fetch
    });
</script>

<div>
    {#if error}
        <p class="error">{error}</p>
    {:else if displayedCompetences}
        <h2>Mes Compétences</h2>
        <div class="competences-grid">
            {#each displayedCompetences as competence}
                <Card
                    competence={competence}
                />
            {/each}
        </div>
    {:else}
        Chargement des compétences...
    {/if}
</div>

<style>
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    .competences-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
        gap: 1em;
        justify-items: center;
        padding: 1em;
    }
</style>