<svelte:options customElement="blog-detail" />

<script>
    import Header from '../elements/header.svelte';
    import { onMount } from 'svelte';
    import { articles, loadArticles } from './store.js';

    let article = null;
    let error = null;

    onMount(async () => {
        const urlParams = new URLSearchParams(window.location.search);
        const id = parseInt(urlParams.get('id'));
        
        if (id) {
            try {
                // Ensure articles are loaded in the store
                await loadArticles();
                
                // Find the article in the store
                article = $articles.find(a => a.ID === id || a.id === id);
                
                if (!article) {
                    error = "Article non trouvé.";
                }
            } catch (err) {
                console.error('Error finding article:', err);
                error = "Erreur de la récupération de l'article.";
            }
        } else {
            error = "ID d'article manquant.";
        }
    });
</script>

<Header/>

<div class="detail-container">
    {#if error}
        <div class="alert alert-danger">{error}</div>
        <div class="flex" style="justify-content: center;">
            <a href="/blog/blog.html" class="btn btn-secondary" style="text-align: center;">
                <i class="bi bi-arrow-left"></i> Retour au blog
            </a>
        </div>
        
    {:else if article}
        <article class="blog-article">
            <header class="article-header">
                <a href="/blog/blog.html" class="back-link"><i class="bi bi-arrow-left"></i> Retour au blog</a>
                <h1>{article.titre}</h1>
                <div class="article-meta">
                    <span class="date"><i class="bi bi-calendar3"></i> {article.date}</span>
                </div>
            </header>
            <hr>
            <section class="article-content">
                {@html article.template}
            </section>
        </article>
    {/if}
</div>

<style>
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");
    @import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css");

    .detail-container {
        max-width: 800px;
        margin: 0 auto;
        padding: 2rem 1rem;
        position: relative;
    }

    @media (min-width: 800px) {
        .detail-container {
            top: var(--header-height);
        }
    }

    .blog-article {
        background: white;
        padding: 2rem;
        border-radius: 8px;
        box-shadow: 0 4px 12px rgba(0,0,0,0.1);
    }

    .article-header h1 {
        padding: 0.5em 0;
        margin: 0;
        color: #333;
    }

    .article-meta {
        color: #999;
        font-size: 0.9rem;
        margin-bottom: 1rem;
    }

    .back-link {
        display: inline-flex;
        align-items: center;
        gap: 0.5rem;
        color: var(--primary-color, #81c553);
        text-decoration: none;
        font-weight: bold;
        margin-bottom: 1rem;
    }

    .back-link:hover {
        text-decoration: underline;
    }

    .article-content {
        line-height: 1.6;
        color: #444;
    }

    .article-content :global(p) {
        margin-bottom: 1.5rem;
    }

    .article-content :global(ul) {
        margin-bottom: 1.5rem;
    }
</style>