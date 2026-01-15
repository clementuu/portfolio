package tmpl

// SpringComp contient la description HTML de la compétence Spring.
const SpringComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        Spring est un framework open source pour le développement d'applications Java. 
        Il fournit une infrastructure complète pour développer des applications Java robustes et performantes, 
        en particulier pour le web. Le framework Spring simplifie le développement Java en fournissant une 
        gamme de fonctionnalités telles que l'injection de dépendances, la gestion des transactions et la sécurité.
    </p>
    <p>
        Mon expérience avec Spring se concentre sur Spring Boot, qui simplifie encore davantage la création d'applications Spring autonomes. 
        J'ai utilisé Spring Boot pour créer des API RESTful, gérer l'accès aux bases de données avec Spring Data, 
        et sécuriser des applications.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Spring, et plus particulièrement <strong>Spring Boot</strong>, est mon framework de prédilection pour le développement d'applications backend Java robustes et performantes. Le projet <strong>Project Management Tool (PMT)</strong> en est une illustration parfaite.
    </p>
    <ul>
        <li>
            <strong>Développement rapide d'API RESTful :</strong> Grâce à Spring Boot, j'ai pu construire efficacement une API RESTful complète et sécurisée pour le PMT. Le framework a grandement facilité la création des endpoints, la gestion des requêtes HTTP et des réponses JSON, accélérant ainsi le développement du backend.
        </li>
        <li>
            <strong>Architecture modulaire et maintenable :</strong> Spring a permis de structurer le code du backend selon une <strong>architecture en couches</strong> (Contrôleur, Service, Repository) en s'appuyant sur les principes d'Inversion de Contrôle (IoC) et d'Injection de Dépendances (DI). Cela garantit une forte cohésion et un faible couplage, essentiels pour la maintenabilité et l'évolutivité.
        </li>
        <li>
            <strong>Gestion de la persistance des données :</strong> J'ai utilisé l'écosystème Spring (notamment <strong>Spring Data JPA</strong>) pour interagir de manière fluide avec la base de données SQL. Cela a simplifié la gestion des entités, l'exécution des requêtes et la manipulation des données, tout en bénéficiant des fonctionnalités avancées de gestion des transactions.
        </li>
        <li>
            <strong>Intégration dans un écosystème Full-Stack :</strong> Spring Boot a joué un rôle clé dans l'intégration harmonieuse du backend Java au sein de l'architecture full-stack du PMT, facilitant sa communication avec le frontend Angular et son déploiement conteneurisé via Docker.
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
