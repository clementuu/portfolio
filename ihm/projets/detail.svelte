<svelte:options customElement="projet-detail" />

<script>
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
                competence = null;
            })
    } else {
        error = "Aucun ID de projet n'a été fourni.";
    }
</script>

<div class="projet-detail-container">
    {#if error}
        <p class="error">{error}</p>
    {:else if projet}
        <h1>{projet.name} - {projet.sujet}</h1>
        {@html projet.tmpl}
    {:else}
        <p>Chargement du projet...</p>
    {/if}
</div>

<style>
    .projet-detail-container {
        padding: 2rem;
        max-width: 900px;
        margin: 0 auto;
        background-color: #fff;
        border-radius: 8px;
        box-shadow: 0 4px 8px rgba(0,0,0,0.1);
    }
    h1 {
        color: #333;
        border-bottom: 2px solid #eee;
        padding-bottom: 2rem;
        margin-bottom: 2rem;
    }
    p {
        color: #555;
        line-height: 1.6;
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

    /*
        * Styles for competence tags within injected HTML
    */
    :global(.competence-tag) {
        display: inline-block;
        padding: 0.2em 0.6em; /* Similar to card-span */
        margin: 0.25em;
        border-radius: 4px; /* Similar to card-span */
        color: white;
        text-decoration: none;
        transition: background-color 0.2s ease;
    }

    :global(.competence-tag.technique) {
        background-color: darkgrey;
    }

    :global(.competence-tag.humain) {
        background-color: darkturquoise;
    }

    :global(.competence-tag.devops) {
        background-color: orchid;
    }

    :global(.competence-tag:hover) {
        filter: brightness(1.1); /* Slightly brighter on hover */
    }
</style>