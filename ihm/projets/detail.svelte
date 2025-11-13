<svelte:options customElement="projet-detail" />

<script>
    import { onMount } from 'svelte';

    let projet = null;
    let error = null;

    onMount(async () => {
        const urlParams = new URLSearchParams(window.location.search);
        const id = urlParams.get('id');

        if (!id) {
            error = "Aucun ID de projet n'a été fourni.";
            return;
        }

        try {
            const response = await fetch(`/projet/${id}`);
            if (!response.ok) {
                throw new Error('La requête a échoué.');
            }
            projet = await response.json();
        } catch (e) {
            error = e.message;
        }
    });
</script>

<style>
    .projet-detail-container {
        padding: 2rem;
        max-width: 900px;
        margin: 2rem auto;
        background-color: #fff;
        border-radius: 8px;
        box-shadow: 0 4px 8px rgba(0,0,0,0.1);
    }
    h1 {
        color: #333;
        border-bottom: 2px solid #eee;
        padding-bottom: 0.5rem;
        margin-bottom: 1rem;
    }
    p {
        color: #555;
        line-height: 1.6;
    }
</style>

<div class="projet-detail-container">
    {#if error}
        <p class="error">{error}</p>
    {:else if projet}
        <h1>{projet.name}</h1>
        <p>{projet.description}</p>
        <!-- Vous pouvez ajouter d'autres détails du projet ici -->
    {:else}
        <p>Chargement du projet...</p>
    {/if}
</div>
