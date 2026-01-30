package tmpl

// SvelteCompetence contient la description HTML de la compétence Svelte.
const SvelteCompetence = `
<section class="comp-section">
    <h3><b>Définition</b></h3>
    <p>
        Svelte est un framework JavaScript moderne qui déplace le travail de l'exécution (runtime) 
        vers la compilation. Contrairement aux frameworks traditionnels comme React ou Vue, Svelte 
        compile le code en JavaScript impératif et optimisé au moment du build, ce qui permet 
        d'obtenir des applications extrêmement rapides avec une taille de bundle réduite.
    </p>
    <p>
        J'ai choisi Svelte pour le développement de l'interface de ce portfolio en raison de sa 
        simplicité, de ses performances et de son approche élégante de la réactivité. Il permet de 
        créer des composants web avec une syntaxe claire et concise, tout en offrant des 
        fonctionnalités puissantes comme le data-binding bidirectionnel et les animations intégrées. C'est un outil très pertinent pour des applications web performantes et légères, idéal pour une expérience utilisateur fluide.
    </p>
</section>

<section class="comp-section">
    <h3><b>Éléments de preuve</b></h3>
    <p>
        J'ai adopté Svelte comme mon framework de prédilection pour le développement d'interfaces web en raison de son approche novatrice (compilateur) et de l'excellente expérience de développement qu'il procure.
    </p>
    
    <p>
        <b>Refonte de modules métier complexes (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>).</b>
        <br>Dans le cadre du projet Escarcelle, j'ai utilisé Svelte pour refondre des modules historiques de l'application de gestion métier. 
        J'ai développé une architecture front-end modulaire basée sur des composants, permettant de gérer des données complexes (produits, stocks) et de 
        communiquer de manière asynchrone avec une API REST. Cette refonte a nettement amélioré la maintenabilité du code et rendu l'interface plus réactive. 
        Ma contribution a consisté à remettre à niveau un pan essentiel du système, en renforçant ses performances et sa capacité d'évolution.
    </p>

    <p>
        <b>Construction d'un portfolio web complet et performant.</b>
        <br>L'intégralité de ce portfolio a été construite avec Svelte. Ce projet sert de démonstration concrète de ma capacité à structurer une application web complète, 
        à créer des composants réutilisables (comme les cartes de projet ou de compétence), à gérer la navigation et à assurer une expérience utilisateur fluide et performante. 
        Le site obtenu est rapide, léger et interactif. J'ai apporté une véritable valeur en créant une vitrine professionnelle soignée, 
        pensée pour mettre en avant mes compétences techniques à travers un outil moderne et performant.
    </p>
</section>

<section class="comp-section">
    <h3><b>Autocritique</b></h3>
    <p>
        Svelte est un des framework frontend que je maîtrise le mieux. Je suis à l'aise avec la réactivité, le système de composants, les stores et les transitions. 
        J'apprécie particulièrement sa philosophie «compiler-first», qui permet des performances proches du natif tout en gardant un framework léger et simple à utiliser. 
        J'ai acquis cette compétence de manière autodidacte et je l'ai approfondie grâce à des projets concrets, mais je sais qu'il me reste encore beaucoup à apprendre pour atteindre un niveau expert.
    </p>
    <p>
        Pour ceux qui découvrent Svelte, je conseillerais de bien comprendre la réactivité et le fonctionnement des stores, 
        et d'expérimenter avec la compilation pour optimiser le rendu. Svelte offre une approche agréable et efficace pour créer des interfaces réactives, 
        et c'est un framework que je trouve particulièrement stimulant.
    </p>
</section>

<section class="comp-section">
    <h3><b>Évolution</b></h3>
    <p>
        À moyen terme, je souhaite approfondir mon usage de SvelteKit pour construire des applications full-stack et mieux maîtriser le Server-Side Rendering (SSR) ainsi que le Static Site Generation (SSG). 
        Je veux également explorer davantage l'écosystème Svelte, notamment les solutions avancées de gestion d'état et les bibliothèques d'UI, 
        afin de créer des interfaces toujours plus performantes et modulaires.
    </p>
    <p>
        Je continue à suivre des tutoriels et à expérimenter avec des projets personnels pour progresser concrètement et, pourquoi pas, me spécialiser dans ce framework à l'avenir.
    </p>
</section>

<section class="comp-section">
    <h5>Projets Associés</h5>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>       
        </ul>
    </div>
</section>
`
