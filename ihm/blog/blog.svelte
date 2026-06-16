<svelte:options customElement="blog-portfolio" />

<script>
    import Header from '../elements/header.svelte';
    import Card from './card.svelte';
    import { onMount } from 'svelte';
    import { articles, loadArticles } from './store.js';

    let error = null;

    onMount(async () => {
        try {
            await loadArticles();
        } catch (err) {
            error = "Erreur de la récupération des articles.";
        }
    });
</script>

<Header/>

<div class="blog-div">
    {#if error}
        <p class="error">{error}</p>
    {:else}
        <h1><b>Bienvenue sur mon blog !</b></h1>
        <div class="blog-feed">
            {#each $articles as article}
                <Card {article} />
            {/each}
        </div>
    {/if}
</div>

<style>
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");
    
    .blog-feed {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 1em 2em;
        max-width: 1000px;
        margin: 0 auto;
    }

    .blog-div {
        position: relative;
        top: 0;
    }

    @media (min-width: 800px) {
		.blog-div {
			top: var(--header-height);
		}
	}
</style>