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


    /* 
        * Les styles ci-dessous utilisent :global() pour pouvoir s'appliquer 
        * au contenu HTML injecté via la directive {@html}.
    */
    :global(.comp-section) {
        margin-bottom: 2rem;
        padding-bottom: 2rem;
        border-bottom: 1px solid #f0f0f0;
    }
    :global(.comp-section:last-of-type) {
        border-bottom: none;
        margin-bottom: 0;
        padding-bottom: 0;
    }

    :global(.comp-section h2) {
        font-size: 1.8em;
        margin-bottom: 1rem;
    }

    :global(.comp-section h3) {
        font-size: 1.4em;
        margin-bottom: 1rem;
        margin-top: 1.5rem;
    }

    :global(.project-list) {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: 1.5rem;
        background-color: #f8f9fa;
        padding: 1.5rem;
        border-radius: 8px;
        border: 1px solid #e9ecef;
        transition: transform 0.2s ease, box-shadow 0.2s ease;
    }

    :global(.project-list strong) {
        display: block;
        font-size: 1.1em;
        color: #343a40;
        margin-bottom: 0.5rem;
    }

    :global(.project-list ul) {
        list-style: none;
        padding: 0;
        margin: 0;
    }

    :global(.project-list li) {
        margin-bottom: 0.5rem;
    }

    :global(a.project-link) {
        text-decoration: none;
        color: #007bff;
        font-weight: 500;
        transition: color 0.2s ease-in-out, text-decoration 0.2s ease-in-out;
    }

    :global(a.project-link:hover) {
        color: #0056b3;
        text-decoration: underline;
    }
</style>