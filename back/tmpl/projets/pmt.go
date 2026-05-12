package tmpl

const PmtHTML = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Présentation</b></h3>

    <p>
        <b>PMT (Project Management Tool)</b> est une application web full-stack développée dans le 
        cadre d'une étude de cas durant mon master. L'objectif était de concevoir un outil complet 
        de gestion de projets permettant de créer des projets, d'organiser des tâches, de les 
        assigner à des membres et de suivre leur avancement via un tableau de bord clair et 
        intuitif.
    </p>

    <img class="project-image" src="../assets/PmtProjet.png" alt="Projet PMT">

    <p>
        L'application repose sur une architecture découplée : un backend en 
        <a class="link" href="/competences/detail.html?id=7">Java/Spring</a> exposant une API REST, 
        et un frontend en <a class="link" href="/competences/detail.html?id=5">Angular</a> consommant 
        cette API. L'ensemble est conteneurisé avec 
        <a class="link" href="/competences/detail.html?id=11">Docker</a>, garantissant un 
        environnement de développement reproductible et cohérent.
    </p>
</section>

<section class="project-section">
    <h3><b>Contexte et Enjeux</b></h3>
    <p>
        Ce projet a été réalisé dans le cadre d'une étude de cas durant mon master, avec pour 
        objectif principal de mettre en pratique une stack technique que je n'avais encore jamais 
        utilisée en conditions réelles.
    </p>

    <p>
        Les enjeux étaient donc avant tout pédagogiques : comprendre les bonnes pratiques d'une 
        architecture full-stack moderne, structurer une API REST cohérente, concevoir une interface 
        utilisateur dynamique et mettre en place un environnement de développement isolé et 
        reproductible.  
        Le projet devait également démontrer ma capacité à mener un développement complet en 
        autonomie, depuis la conception jusqu'au déploiement.
    </p>

    <p>
        PMT a ainsi servi de terrain d'expérimentation pour consolider mes compétences en 
        développement web, en architecture logicielle et en outillage DevOps.
    </p>

    <h3><b>Acteurs et Interactions</b></h3>
    <p>
        PMT est un projet entièrement personnel, conçu et développé en autonomie. J'ai assuré 
        l'ensemble des rôles nécessaires à sa réalisation :
    </p>

    <ul>
        <li>
            <strong>Moi-même</strong> - responsable de la conception 
            fonctionnelle, du développement backend et frontend, des tests, de la conteneurisation 
            et utilisateur de l'outil.
        </li>
    </ul>

    <p>
        Cette configuration m'a permis d'expérimenter un cycle complet de développement logiciel, 
        depuis l'idéation jusqu'à l'usage réel, en passant par la conception technique, les choix 
        d'architecture et la mise en place d'un environnement de déploiement reproductible.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISSIONS ET COMPÉTENCES DÉVELOPPÉES                       -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Missions et Compétences Développées</b></h3>
    <p>
        Ce projet complet m'a permis de maîtriser l'ensemble du cycle de vie d'une application web 
        moderne, de la conception de la base de données à la conteneurisation, en passant par le 
        développement backend et frontend.
    </p>

    <div class="skills-grid">

        <div class="skill-item">
            <strong>Conception & Architecture</strong>
            <p>
                Élaboration du modèle de données, définition des entités et des relations.  
                Conception de l'architecture globale (API REST, SPA Angular, conteneurs Docker).  
                Rédaction des spécifications fonctionnelles et techniques.
            </p>
        </div>

        <div class="skill-item">
            <strong>Développement Backend (Spring)</strong>
            <p>
                Mise en place d'une API RESTful robuste avec Spring Boot.  
                Structuration en couches (Controller, Service, Repository) pour une meilleure 
                maintenabilité.  
                Gestion de la sécurité, de la validation des données et de la configuration.
            </p>
        </div>

        <div class="skill-item">
            <strong>Développement Frontend (Angular)</strong>
            <p>
                Création d'une application monopage (SPA) dynamique.  
                Développement de composants modulaires (Dashboard, Project, Task).  
                Gestion de l'état et communication asynchrone avec le backend via des services.
            </p>
        </div>

        <div class="skill-item">
            <strong>Architecture Full-Stack & Intégration</strong>
            <p>
                Conception du contrat d'API, gestion des requêtes HTTP, authentification et règles CORS.  
                Mise en place d'une communication fluide entre le client et le serveur.
            </p>
        </div>

        <div class="skill-item">
            <strong>Gestion de Base de Données (SQL)</strong>
            <p>
                Modélisation du schéma relationnel.  
                Création des entités (Projet, Tâche, Utilisateur) et gestion des relations.  
                Scripts d'initialisation pour faciliter les tests.
            </p>
        </div>

        <div class="skill-item">
            <strong>DevOps & Conteneurisation (Docker)</strong>
            <p>
                Création d'un environnement de développement isolé avec Docker et Docker Compose.  
                Conteneurisation du backend, du frontend et de la base de données.  
                Configuration d'un réseau interne pour orchestrer les services.
            </p>
        </div>

        <div class="skill-item">
            <strong>Automatisation & Outillage (Maven, Makefile)</strong>
            <p>
                Utilisation de Maven pour la gestion des dépendances et le build Java.  
                Création d'un Makefile pour automatiser les tâches courantes (build, tests, lancement).
            </p>
        </div>

    </div>
</section>

<!-- =================================================================== -->
<!-- SECTION : CONCLUSION                                                -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Résultats</b></h3>
    <p>
        PMT est un projet intégrateur qui démontre ma capacité à concevoir, développer et déployer 
        une application web complète en autonomie.  
        Il met en évidence ma maîtrise de la stack <b>Java/Spring</b> côté serveur, 
        <b>Angular</b> côté client, ainsi que des outils DevOps modernes comme <b>Docker</b>.  
        Ce projet a été un excellent exercice pour comprendre les enjeux d'une architecture 
        full-stack moderne et renforcer ma polyvalence.
    </p>

    <h3><b>Pour le futur</b></h3>
    <p>
        PMT est destiné à devenir mon outil personnel pour organiser mes projets.  
        J'envisage de le rendre accessible en ligne, ce qui nécessitera une gestion multi-utilisateurs, 
        une authentification renforcée et une réarchitecture partielle pour supporter plusieurs 
        organisations.  
        Ces évolutions permettront d'en faire un outil plus générique, plus sécurisé et potentiellement 
        utilisable par un public plus large.
    </p>

    <h3><b>Mon analyse critique</b></h3>
    <p>
        Avec le recul, PMT a été un projet très formateur. Il m'a permis de découvrir une stack 
        technique complète et de comprendre les bonnes pratiques d'un développement full-stack.  
        J'aurais pu aller plus loin sur certains aspects, notamment les tests automatisés, la gestion 
        fine des rôles utilisateurs ou la CI/CD.  
        Malgré cela, ce projet reste une excellente démonstration de ma capacité à mener un 
        développement complet en autonomie.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : COMPÉTENCES ASSOCIÉES                                     -->
<!-- =================================================================== -->
<section class="project-section">
    <h5>Compétences Associées</h5>
    <div class="competences-list">
        <a href="/competences/detail.html?id=7" class="competence-tag technique">Spring</a>
        <a href="/competences/detail.html?id=5" class="competence-tag technique">Angular</a>
        <a href="/competences/detail.html?id=6" class="competence-tag technique">SQL</a>
        <a href="/competences/detail.html?id=11" class="competence-tag devops">Docker</a>
        <a href="/competences/detail.html?id=15" class="competence-tag humain">Créativité</a>
        <a href="/competences/detail.html?id=16" class="competence-tag humain">Esprit critique</a>
    </div>
</section>
`
