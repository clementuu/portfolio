<svelte:options customElement="competences-portfolio" />

<script>
    import Card from './card.svelte';

    let competences = [];

    async function getCompetences() {
        try {
            const response = await fetch('/competences');

            if (!response.ok) {
                // Gestion des erreurs HTTP (ex: 404, 500)
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }

            const data = await response.json();
            competences = data;
        } catch (error) {
            // Affichage plus clair et utile pour le développeur
            console.error("Échec de la récupération des compétences :", error.message || error);
        }
    }

    getCompetences();
</script>

<div>
    <h2>Mes Compétences</h2>
    <div class="competences-grid">
        {#each competences as competence}
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