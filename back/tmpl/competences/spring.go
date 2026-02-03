package tmpl

// SpringCompetence contient la description HTML de la compétence Spring.
const SpringCompetence = `
<section class="comp-section">
    <h3><b>Définition</b></h3>
    <p>
        Spring est un framework open source majeur de l'écosystème Java, conçu pour faciliter le développement 
        d'applications robustes, modulaires et sécurisées. Dans un contexte professionnel où les entreprises 
        recherchent des solutions fiables, scalables et compatibles avec les architectures cloud-native, 
        Spring s'impose comme un standard incontournable. L'actualité technologique - montée en puissance 
        des microservices, généralisation des API REST, adoption massive de Kubernetes - renforce encore 
        son importance dans les systèmes d'information modernes.
    </p>
    <p>
        J'ai découvert Java dans le cadre universitaire. C'est le tout premier langage compilé auquel j'ai été confronté, 
        avant que je ne tombe réellement dans l'informatique. Ça a pu être frustrant au début et ce n'est pas l'amour fou, 
        mais sans lui je ne me serais peut-être moins intéressé à la programmation. C'est réellement avec Spring que j'ai accroché à l'univers Java. 
        Je l'utilise pour concevoir des API performantes, gérer la persistance des données via Spring Data, structurer des architectures en couches 
        maintenables et intégrer des mécanismes de sécurité adaptés aux besoins métier.
    </p>
</section>

<section class="comp-section">
    <h3><b>Éléments de preuve</b></h3>
    <p>
        Ma maîtrise de Spring Boot s'est construite à travers des projets concrets où j'ai dû concevoir, 
        optimiser et sécuriser des architectures backend complètes. Voici trois anecdotes illustrant 
        l'application directe de cette compétence.
    </p>

    <p>
        <strong>Développement d'une API REST complète et sécurisée (<a class="link" href="/projets/detail.html?id=6">PMT</a>).</strong>
        <br>Dans le cadre du projet Project Management Tool (PMT), j'ai conçu une API RESTful 
        permettant la gestion des projets, des tâches et des utilisateurs. Spring Boot m'a permis de 
        structurer rapidement les endpoints et de gérer les flux JSON facilement. 
        En respectant les convention et l'esprit du framework, on obtient un backend fiable, performant et extensible, 
        facilitant l'évolution du produit.  
    </p>

    <p>
        <strong>Architecture modulaire et maintenable (<a class="link" href="/projets/detail.html?id=6">PMT</a>).</strong>
        <br>Pour garantir la qualité du code et la facilité d'évolution, j'ai structuré le backend selon une 
        architecture en couches (Controller, Service, Repository) en appliquant les principes d'Inversion 
        de Contrôle (IoC) et d'Injection de Dépendances (DI). Résultat : un code propre, testable et durable, 
        facilitant l'intégration de nouvelles fonctionnalités et permettant à d'autres développeurs de contribuer efficacement. 
    </p>

    <p>
        <strong>Intégration dans un environnement Full-Stack conteneurisé (<a class="link" href="/projets/detail.html?id=6">PMT</a>).</strong>
        <br>Le service Java/Spring de PMT s'intégrait parfaitement dans un écosystème moderne : il dialoguait avec une base de données SQL pour la persistance des données, servait de socle pour le front-end 
        <a class="link" href=/competences/detail.html?id=5>Angular</a>, et l'ensemble était orchestré et déployé dans un environnement conteneurisé avec <a class="link" href=/competences/detail.html?id=12>Docker</a>. 
        L'application obtenue est cohérente, fonctionnelle et prête à être déployée. J'ai ainsi assuré une intégration harmonieuse entre toutes les couches, garantissant une architecture stable et homogène.
    </p>
</section>

<section class="comp-section">
    <h3><b>Autocritique</b></h3>
    <p>
        Ma maîtrise de Spring est opérationnelle. J'ai une bonne compréhension des concepts orientés objet, du framework et des concepts de base de Java. 
        J'ai découvert le framework dans le cadre de mes études puis à travers un projet personnel, 
        et c'est ce qui m'a réellement réconcilié avec l'écosystème Java. 
        Ces expériences m'ont permis de comprendre les mécanismes essentiels du framework et de développer des API structurées, 
        maintenables et adaptées aux besoins métier.
    </p>
    <p>
        Même si mon expérience reste limitée en volume, elle m'a donné une bonne compréhension des fondamentaux: 
        injection de dépendances, gestion des entités, transactions, architecture en couches. 
        Je suis convaincu que la solidité d'un backend repose d'abord sur ces bases avant d'aborder des architectures plus ambitieuses.
    </p>
</section>

<section class="comp-section">
    <h3><b>Évolution</b></h3>
    <p>
        À moyen terme, je souhaite approfondir mes compétences autour de l'écosystème Spring, notamment 
        sur des sujets avancés comme Spring Security, et l'intégration avec des orchestrateurs tels que Kubernetes. 
        Mon objectif est de maîtriser la conception de microservices distribués, résilients et observables.
    </p>
</section>

<section class="comp-section">
    <h5>Projets Associés</h5>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
