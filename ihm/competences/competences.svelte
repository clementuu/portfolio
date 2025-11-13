<svelte:options customElement="competences-portfolio" />

<script>
    import Card from './card.svelte';
    import { onMount } from 'svelte'; // Import onMount

    let allCompetences = []; // Store all fetched competencies
    let displayedCompetences = []; // Competencies to display after filtering

    onMount(async () => { // Use onMount to fetch data
        await getCompetences();
        filterCompetences(); // Filter after fetching
    });

    async function getCompetences() {
        try {
            const response = await fetch('/competences');

            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }

            const data = await response.json();
            allCompetences = data; // Store all competencies
        } catch (error) {
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
</script>

<div>
    <h2>Mes Compétences</h2>
    <div class="competences-grid">
        {#each displayedCompetences as competence}
            <Card
                competence={competence}
            />
        {/each}
    </div>
</div>

<style>
	@import '../style/style.css';
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    .competences-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
        gap: 1em;
        justify-items: center;
        padding: 1em;
    }
</style>