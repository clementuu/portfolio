package tmpl

// SvelteComp contient la description HTML de la compétence Svelte.
const SvelteComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
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
        fonctionnalités puissantes comme le data-binding bidirectionnel et les animations intégrées.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        J'ai adopté Svelte comme mon framework de prédilection pour le développement d'interfaces web en raison de son approche novatrice (compilateur) et de l'excellente expérience de développement qu'il procure. Mon expertise s'illustre à travers deux projets principaux :
    </p>
    <ul>
        <li>
            <strong>Application de gestion métier (Escarcelle) :</strong> Svelte a été un levier majeur dans la modernisation du projet Escarcelle. Je l'ai utilisé pour refondre des modules historiques, en développant une <strong>architecture front-end modulaire</strong> basée sur des composants. Ce travail a permis d'améliorer drastiquement la maintenabilité du code tout en construisant des interfaces de gestion de données complexes (produits, stocks, etc.) qui communiquent de manière asynchrone avec une API REST.
        </li>
        <li>
            <strong>Portfolio personnel (ce site web) :</strong> L'intégralité de ce portfolio a été construite avec Svelte. Il sert de démonstration concrète de ma capacité à structurer une application web complète, à créer des <strong>composants réutilisables</strong> (comme les cartes de projet ou de compétence), à gérer la navigation et à assurer une expérience utilisateur fluide et performante.
        </li>
    </ul>
    <p>
        Dans ces projets, Svelte me permet de produire des applications rapides et légères, en écrivant un code concis et lisible, tout en bénéficiant d'un système de réactivité puissant et intuitif.
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
