<svelte:options customElement="competences-portfolio" />

<script>
    import Header from '../elements/header.svelte';
    import Card from './card.svelte';

    let error = null;
    let allCompetences = [];
    let displayedCompetences = null; // compétences à afficher après filtrage

    fetch('/competences')
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            allCompetences = data;
            filterCompetences(); // filtre les données après le fetch
        })
        .catch(e => {
            console.error("Error fetching competences:",e);
            error = "Erreur lors du chargement des compétences."
            allCompetences = [];
            displayedCompetences = null;
        });

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

<Header/>

<div class="competences-div">
    {#if error}
        <p class="error">{error}</p>
    {:else if displayedCompetences}
        <h1><b>Mes Compétences</b></h1>
        <div class="competences-grid">
            {#each displayedCompetences as competence}
                <Card
                    competence={competence}
                />
            {/each}
        </div>
    {/if}
</div>

<style>
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    .competences-grid {
        display: inline-flex;
        flex-wrap: wrap;
        justify-content: center;
        padding: 1em;
    }

    .competences-div {
        position: relative;
        top: 0;
    }

    @media (min-width: 800px) {
		.competences-div {
			top: var(--header-height);
		}
	}
</style>