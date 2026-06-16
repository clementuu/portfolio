<svelte:options customElement="blog-detail" />

<script>
    import Header from '../elements/header.svelte';
    import { onMount } from 'svelte';

    let article = null;
    let error = null;

    // Mock data (same as in blog.svelte for now)
    const articles = [
        {
            ID: 1,
            titre: "Mon premier article",
            template: `
                <p>Bienvenue sur mon blog !</p>
                <p>Dans cet article, je partage ma vision du développement logiciel et pourquoi j'ai décidé de créer ce portfolio.</p>
                <p>L'ingénierie logicielle est pour moi un outil au service de l'humain. C'est pourquoi j'ai choisi de mettre en avant mon parcours et mes projets.</p>
                <p>J'espère que vous apprécierez la lecture !</p>
            `,
            date: "01/01/2026"
        },
        {
            ID: 2,
            titre: "Pourquoi Go ?",
            template: `
                <p>Go est devenu mon langage de prédilection pour le backend.</p>
                <p>Sa simplicité, sa performance et son excellent support pour la concurrence en font un choix idéal pour les microservices modernes.</p>
                <ul>
                    <li>Simplicité de lecture et d'écriture</li>
                    <li>Performance proche du C++</li>
                    <li>Gestion native de la concurrence avec les goroutines</li>
                </ul>
            `,
            date: "05/01/2026"
        },
        {
            ID: 3,
            titre: "Svelte : la simplicité au front",
            template: `
                <p>Après avoir utilisé Angular et React, j'ai découvert Svelte.</p>
                <p>Sa philosophie 'sans framework' à l'exécution et sa syntaxe concise m'ont immédiatement séduit.</p>
                <p>Svelte compile votre code en JavaScript pur, ce qui signifie pas de virtual DOM et des performances accrues.</p>
            `,
            date: "12/03/2026"
        }
    ];

    onMount(() => {
        const urlParams = new URLSearchParams(window.location.search);
        const id = parseInt(urlParams.get('id'));
        
        if (id) {
            article = articles.find(a => a.ID === id);
            if (!article) {
                error = "Article non trouvé.";
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
        <a href="/blog/blog.html" class="btn btn-secondary">Retour au blog</a>
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