package tmpl

// AngularComp contient la description HTML de la compétence Angular.
const AngularComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        Angular est un framework de développement d'applications web open-source basé sur TypeScript, dirigé par l'équipe Angular de Google. 
        Il est conçu pour créer des applications web monopages (Single Page Applications) complexes et performantes. 
        Angular suit une architecture basée sur les composants et fournit une suite complète d'outils pour le routage, la gestion d'état et la communication avec les API.
    </p>
    <p>
        Mon expérience avec Angular comprend la création de composants, la gestion des services et de l'injection de dépendances, 
        et l'utilisation du routage pour créer une expérience de navigation fluide.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Ma compétence avec Angular a été mise en œuvre de manière concrète lors du développement du <strong>Project Management Tool (PMT)</strong>, une application web full-stack où Angular constituait la pierre angulaire du front-end.
    </p>
    <ul>
        <li>
            <strong>Architecture d'une Single-Page Application (SPA) :</strong> J'ai conçu et développé l'intégralité de l'application front-end comme une SPA dynamique. Cela a impliqué la mise en place d'une <strong>architecture à base de composants</strong> (Dashboard, Projet, Tâche) pour créer une interface modulaire, maintenable et évolutive.
        </li>
        <li>
            <strong>Gestion de l'état et services :</strong> Pour gérer la logique métier et les données, j'ai utilisé des services Angular afin d'encapsuler la communication asynchrone avec l'API backend. Cette approche a permis de centraliser la gestion de l'état de l'application et de découpler les composants de la logique de récupération des données.
        </li>
        <li>
            <strong>Intégration dans un écosystème Full-Stack :</strong> Le front-end Angular a été développé pour consommer une <strong>API RESTful</strong> construite en Java et Spring Boot. Cette expérience m'a permis de maîtriser le cycle complet de l'intégration, de la définition du contrat de données à la gestion des requêtes HTTP et des réponses, le tout dans un environnement conteneurisé avec Docker.
        </li>
    </ul>
</section>
<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
