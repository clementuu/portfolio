<svelte:options customElement="projets-portfolio" />

<script>
    import Card from './card.svelte';
    import { onMount } from 'svelte';

    let projets = [];

    async function getProjets() {
        try {
            const response = await fetch('/projets');

            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }

            const data = await response.json();
            projets = data; 
        } catch (error) {
            console.error("Échec de la récupération des compétences :", error.message || error);
        }
    }

    onMount(async () => {
        await getProjets();
        console.log(projets);
    });
</script>

<div>
    <h2>Mes Projets</h2>
    <div class="projets-grid">
        {#each projets as projet}
            <Card {projet} />
        {/each}
    </div>
</div>

<style>
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");
  
    .projets-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
        gap: 1em;
        justify-items: center;
        padding: 1em;
    }
</style>