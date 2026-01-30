<svelte:options customElement="projet-detail" />

<script>
    import Header from "../elements/header.svelte";

    let projet = null;
    let error = null;

    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    if (id) {
        fetch(`/projet/${id}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                projet = data;
                error = null;
            })
            .catch(e => {
                console.error("Error fetching projet:", e);
                error = "Erreur lors du chargement des détails du projet.";
                projet = null;
            })
    } else {
        error = "Aucun ID de projet n'a été fourni.";
    }
</script>

<Header/>

<div class="container projet-detail">
    {#if error}
        <p class="error">{error}</p>
    {:else if projet}
        <h1><b>{projet.name}</b></h1>
        <h2><b>{projet.sujet}</b></h2>
        {@html projet.tmpl}
    {:else}
        <p>Chargement du projet...</p>
    {/if}
</div>

<style>
    @import "../style/style.css";
    
    .projet-detail {
        position: relative;
        top: 0;
        margin: auto;
        padding: 2rem;
        max-width: 900px;
        background-color: #fff;
    }

    h1 {
        padding-bottom: 10px;
        margin-bottom: 0;
    }

    h2 {
        border-bottom: 2px solid #eee;
        padding-top: 0;
        padding-bottom: 2rem;
        padding-left: 0;
        margin-top: 0;
        margin-bottom: 2rem;
    }
    p {
        color: #555;
        line-height: 1.6;
    }

    @media (min-width: 770px) {
		.projet-detail {
			top: var(--header-height);
            margin: 1em auto 5em auto;
		}
	}

    /* 
     * Les styles ci-dessous utilisent :global() pour pouvoir s'appliquer 
     * au contenu HTML injecté via la directive {@html}.
    */
    :global(.project-section) {
        margin-bottom: 2rem;
        padding-bottom: 2rem;
        border-bottom: 1px solid #f0f0f0;
    }

    :global(.project-section:last-of-type) {
        border-bottom: none;
        margin-bottom: 0;
        padding-bottom: 0;
    }

    :global(.project-section h2) {
        font-size: 1.8em;
        margin-bottom: 1rem;
    }

    :global(.project-section h3) {
        font-size: 1.4em;
        margin-bottom: 1rem;
        margin-top: 1.5rem;
    }

    :global(.skills-grid) {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: 1.5rem;
    }

    :global(.skill-item) {
        background-color: #f8f9fa;
        padding: 1.5rem;
        border-radius: 8px;
        border: 1px solid #e9ecef;
        transition: transform 0.2s ease, box-shadow 0.2s ease;
    }

    :global(.skill-item strong) {
        display: block;
        font-size: 1.1em;
        color: #343a40;
        margin-bottom: 0.5rem;
    }

    :global(.competences-list) {
        display: inline-block;
        margin: auto;
        justify-content: center;
    }

    :global(.competence-tag) {
        display: inline-block;
        padding: 0.2em 0.6em; /* Similar to card-span */
        margin: 0.25em;
        border-radius: 4px; /* Similar to card-span */
        color: black;
        text-decoration: none;
        transition: background-color 0.2s ease;
    }

    :global(.competence-tag.technique) {
        background-color: var(--dev-color);
    }

    :global(.competence-tag.humain) {
        background-color: var(--human-color);
    }

    :global(.competence-tag.devops) {
        background-color: var(--devops-color);
    }

    :global(.competence-tag:hover) {
        filter: brightness(1.1); /* Slightly brighter on hover */
    }
</style>