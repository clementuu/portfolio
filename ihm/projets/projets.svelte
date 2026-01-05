<svelte:options customElement="projets-portfolio" />

<script>
    import Card from './card.svelte';

    let projets = [];
    let error = null;
    
    fetch('/projets')
        .then(response => {
            if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
            return response.json();
        })
        .then(data => {
            projets = data;
        })
        .catch(err => {
            console.error('Fetch error:', err);
            error = "Erreur de la récupération des compétences.";
            projets = [];
        });
</script>

<div>
    {#if error}
        <p class="error">{error}</p>
    {:else}
        <h2>Mes Projets</h2>
        <div class="projets-grid">
            {#each projets as projet}
                <Card {projet} />
            {/each}
        </div>
    {/if}
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