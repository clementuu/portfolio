<svelte:options customElement="blog-portfolio" />

<script>
    import Header from '../elements/header.svelte';
    import Card from './card.svelte';

    let articles = [
        {
            ID: 1,
            titre: "Mon premier article",
            description: "Bienvenue sur mon blog ! Dans cet article, je partage ma vision du développement logiciel et pourquoi j'ai décidé de créer ce portfolio. L'ingénierie logicielle est pour moi un outil au service de l'humain.",
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
            description: "Go est devenu mon langage de prédilection pour le backend. Sa simplicité, sa performance et son excellent support pour la concurrence en font un choix idéal pour les microservices modernes.",
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
            description: "Après avoir utilisé Angular et React, j'ai découvert Svelte. Sa philosophie 'sans framework' à l'exécution et sa syntaxe concise m'ont immédiatement séduit.",
            template: `
                <p>Après avoir utilisé Angular et React, j'ai découvert Svelte.</p>
                <p>Sa philosophie 'sans framework' à l'exécution et sa syntaxe concise m'ont immédiatement séduit.</p>
                <p>Svelte compile votre code en JavaScript pur, ce qui signifie pas de virtual DOM et des performances accrues.</p>
            `,
            date: "12/03/2026"
        }
    ];
</script>

<Header/>

<div class="blog-div">
    <h1><b>Bienvenue sur mon blog !</b></h1>
    <div class="blog-feed">
        {#each articles as article}
            <Card {article} />
        {/each}
    </div>
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