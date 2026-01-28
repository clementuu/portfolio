package tmpl

// SvelteComp contient la description HTML de la compétence Svelte.
const SvelteComp = `
<section class="comp-section">
    <h3>Définition</h3>
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
    <h3>Éléments de preuve</h3>
    <p>
        J'ai adopté Svelte comme mon framework de prédilection pour le développement d'interfaces web en raison de son approche novatrice (compilateur) et de l'excellente expérience de développement qu'il procure.
    </p>
    
    <p>
        <strong>Refonte de modules métier complexes (Escarcelle).</strong>
        Dans le cadre du projet <strong>Escarcelle</strong>, j'ai utilisé Svelte pour refondre des modules historiques de l'application de gestion métier. J'ai développé une <strong>architecture front-end modulaire</strong> basée sur des composants, permettant de gérer des données complexes (produits, stocks) et de communiquer de manière asynchrone avec une API REST. Le résultat a été une amélioration drastique de la maintenabilité du code et une interface utilisateur plus réactive. Ma valeur ajoutée a été de moderniser une partie critique du système, rendant l'application plus performante et plus facile à faire évoluer. (Voir <a href="/projets/detail.html?id=1">Escarcelle</a>)
    </p>

    <p>
        <strong>Construction d'un portfolio web complet et performant.</strong>
        L'intégralité de ce <strong>portfolio personnel</strong> a été construite avec Svelte. Ce projet sert de démonstration concrète de ma capacité à structurer une application web complète, à créer des <strong>composants réutilisables</strong> (comme les cartes de projet ou de compétence), à gérer la navigation et à assurer une expérience utilisateur fluide et performante. Le résultat est un site web rapide, léger et interactif. Ma valeur ajoutée a été de concevoir et de réaliser une vitrine professionnelle qui met en lumière mes compétences techniques avec un outil de pointe.
    </p>
</section>

<section class="comp-section">
    <h3>Autocritique</h3>
    <p>
        Maîtrise de Svelte : <strong>Avancée</strong>. Je suis très à l'aise avec la réactivité, le système de composants, les stores et les transitions de Svelte. J'apprécie particulièrement sa philosophie "compiler-first" qui offre des performances natives sans le poids d'un runtime. C'est une compétence que j'ai acquise de manière autodidacte et que j'ai rapidement approfondie grâce à des projets concrets.
    </p>
    <p>
        Je conseillerais aux nouveaux venus de se concentrer sur la compréhension de la réactivité de Svelte et de ne pas hésiter à exploiter les fonctionnalités de compilation pour optimiser leurs applications. Svelte est un excellent choix pour quiconque souhaite développer des interfaces utilisateur performantes avec une courbe d'apprentissage douce.
    </p>
</section>

<section class="comp-section">
    <h3>Évolution</h3>
    <p>
        À moyen terme, je souhaite explorer les possibilités offertes par SvelteKit pour la construction d'applications full-stack et approfondir mes connaissances en Server-Side Rendering (SSR) et Static Site Generation (SSG) avec Svelte. Je prévois également de m'intéresser davantage aux écosystèmes autour de Svelte pour l'intégration avec des outils de gestion d d'état plus complexes et des bibliothèques d'UI avancées.
    </p>
    <p>
        Je suis actuellement des tutoriels sur SvelteKit et je compte développer de petits projets personnels pour expérimenter ces nouvelles facettes.
    </p>
</section>

<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>       
        </ul>
    </div>
</section>
`
