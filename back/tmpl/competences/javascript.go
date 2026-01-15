package tmpl

// JavaScriptComp contient la description HTML de la compétence JavaScript.
const JavaScriptComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        JavaScript est un langage de programmation de scripts principalement employé dans les pages web interactives, 
        mais il est aussi utilisé côté serveur (avec Node.js), pour des applications mobiles et de bureau. 
        C'est un langage dynamique, interprété et basé sur les prototypes.
    </p>
    <p>
        Mon expérience en JavaScript couvre la manipulation du DOM, la gestion d'événements, 
        les appels asynchrones (AJAX/Fetch) et l'utilisation de bibliothèques et de frameworks modernes. 
        Je l'ai utilisé pour créer des expériences utilisateur interactives et dynamiques.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Ma maîtrise de JavaScript s'étend du développement "vanilla" (sans framework) pour des expériences interactives sur-mesure, à son utilisation au cœur d'applications web modernes via des frameworks réactifs.
    </p>
    <ul>
        <li>
            <strong>Développement de jeu 2D en JavaScript natif :</strong> Le projet <strong>Jeu Anti-Gaspi</strong> est l'illustration la plus parlante de cette compétence. Développé sans aucun framework, il m'a permis de mettre en œuvre des mécaniques de jeu complexes : manipulation du <strong>Canvas HTML5</strong> pour le rendu, création d'une boucle de jeu avec <code>requestAnimationFrame</code>, gestion des animations, détection de collisions et gestion des entrées utilisateur.
        </li>
        <li>
            <strong>Socle des applications web modernes :</strong> JavaScript est le langage fondamental sur lequel reposent les frameworks que j'utilise. Dans des projets comme <strong>Escarcelle</strong> ou ce même <strong>portfolio</strong>, mes compétences en JavaScript sont appliquées à travers <strong>Svelte</strong> pour créer des composants d'interface réactifs, gérer l'état de l'application et interagir de manière asynchrone avec les APIs backend.
        </li>
        <li>
            <strong>Manipulation du DOM et interactivité :</strong> Au-delà des frameworks, j'utilise JavaScript pour enrichir les pages web en manipulant directement le DOM, en validant des formulaires côté client et en ajoutant toutes sortes d'interactions pour améliorer l'expérience utilisateur.
        </li>
    </ul>
</section>
<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=4">Jeu Anti-Gaspi</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
        </ul>
    </div>
</section>
`
