<svelte:options customElement="competence-detail" />

<script>
    import { onMount } from 'svelte';

    let competence = null;
    let error = null;

    onMount(() => {
        const storedCompetence = sessionStorage.getItem('selectedCompetence');

        if (storedCompetence) {
            try {
                const competenceData = JSON.parse(storedCompetence);
                competence = competenceData;
            } catch (e) {
                error = "Erreur lors de la lecture des données de la compétence.";
            }
        } else {
            error = "Aucune donnée de compétence trouvée. Veuillez retourner à la page des compétences et en sélectionner une.";
        }
    });

    function getStars(rating) {
        if (rating === null || rating === undefined) return '';
        let stars = '';
        for (let i = 0; i < 5; i++) {
            if (i < rating) {
                stars += '⭐';
            } else {
                stars += '☆';
            }
        }
        return stars;
    }
</script>

<div class="container">
    {#if error}
        <p class="error">{error}</p>
    {:else if competence}
        <div class="detail-container">
            <h1>{competence.name}</h1>
            <img src={competence.image} alt={`Logo de ${competence.name}`} />
            <div class="stars">{getStars(competence.rating)}</div>
            <div class="detail-text">{@html competence.desc}</div>
        </div>
    {:else}
        <p>Chargement des détails de la compétence...</p>
    {/if}
</div>

<style>
    @import "../style/style.css";
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    .detail-container {
        padding: 2em;
        max-width: 800px;
        margin: 2em auto;
        background: #f9f9f9;
        border-radius: 8px;
        box-shadow: 0 4px 12px rgba(0,0,0,0.1);
        text-align: center;
    }
    .detail-container img {
        max-width: 100px;
        margin-bottom: 1em;
    }
    .stars {
        font-size: 1.5em;
        color: gold;
        margin-bottom: 1em;
    }
    .detail-text {
        text-align: justify;
        margin-top: 1.5em;
        line-height: 1.6;
    }
</style>