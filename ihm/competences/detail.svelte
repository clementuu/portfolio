<svelte:options customElement="competence-detail" />

<script>
    let competence = null;
    let error = null;
    let loading = true;

    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    if (id) {
        loading = true;
        fetch(`/competence/${id}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                competence = data;
                error = null;
            })
            .catch(e => {
                console.error("Error fetching competence:", e);
                error = "Erreur lors du chargement des détails de la compétence.";
                competence = null;
            })
            .finally(() => {
                loading = false;
            });
    } else {
        error = "Aucun ID de compétence n'a été fourni.";
    }

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

<div class="container competence-detail">
    {#if error}
        <p class="error">{error}</p>
    {:else if loading}
        <p>Chargement des détails de la compétence...</p>
    {:else if competence}
        <h1>{competence.name}</h1>
        <img src={competence.image} alt={`Logo de ${competence.name}`} />
        <div class="stars">{getStars(competence.rating)}</div>
        <div class="detail-text">{@html competence.template}</div>
    {:else}
        <p>Aucune donnée de compétence trouvée. Veuillez retourner à la page des compétences et en sélectionner une.</p>
    {/if}
</div>

<style>
    @import "../style/style.css";
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    .competence-detail {
        position: relative;
        top: 0;
        margin: auto;
        padding: 2em;
        max-width: 800px;
        text-align: center;
    }

    .competence-detail img {
        max-width: 50px;
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

    @media (min-width: 768px) {
		.competence-detail {
			top: var(--header-height);
            margin: 1em auto 5em auto;
		}
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
        display: flex;
        gap: 1.5rem;
        background-color: #f8f9fa;
        padding: 1.5rem;
        transition: transform 0.2s ease, box-shadow 0.2s ease;
        justify-content: center;
    }

    :global(.project-list strong) {
        display: block;
        font-size: 1.1em;
        color: #343a40;
        margin-bottom: 0.5rem;
    }

    :global(.project-list ul) {
        list-style: none;
        display: flex;
        padding: 0;
        margin: 0;
    }

    :global(.project-list li) {
        margin-bottom: 0.5rem;
    }

    :global(a.project-link) {
        display: inline-block;
        padding: 0.2em 0.6em; /* Similar to card-span */
        margin: 0.25em;
        border-radius: 4px; /* Similar to card-span */
        color: white;
        text-decoration: none;
        transition: background-color 0.2s ease;
        text-decoration: none;
        background-color: #007bff;
        font-weight: 500;
    }

    :global(a.project-link:hover) {
        filter: brightness(1.1); /* Slightly brighter on hover */
    }

    :global(a.link) {
        color: #212529;
        font-weight: bold;
        text-decoration: none;
    }
    :global(a.link:hover) {
        color: #007bff;
    }
</style>