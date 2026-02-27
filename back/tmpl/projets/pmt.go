package tmpl

const PmtHTML = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <p>
        <b>PMT (Project Management Tool)</b> est une application web full-stack conçue pour la gestion 
        de projets. Elle permet aux utilisateurs de créer des projets, de définir des tâches, 
        de les assigner à des membres et de suivre leur avancement au sein d'un tableau de bord 
        intuitif.
    </p>

    <p>
        Le projet est construit sur une architecture découplée avec un backend en 
        <a class="link" href="/competences/detail.html?id=7">Java/Spring</a> exposant une API REST, et un frontend en <a class="link" href="/competences/detail.html?id=5">Angular</a> 
        consommant cette API. L'ensemble est conteneurisé avec <a class="link" href="/competences/detail.html?id=11">Docker</a>, garantissant 
        un environnement de développement et de déploiement cohérent et reproductible.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : ACTEURS ET INTERACTIONS DU PROJET                        -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Acteurs et Interactions</b></h3>
    <p>
        PMT est un projet entièrement personnel, conçu et développé en autonomie. J'ai assuré 
        l'ensemble des rôles nécessaires à sa réalisation :
    </p>

    <ul>
        <li>
            <strong>Moi-même (développeur, testeur et utilisateur)</strong> - responsable de la conception 
            fonctionnelle, du développement backend et frontend, des tests, de la conteneurisation 
            et de l'utilisation quotidienne de l'outil.
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
        Ce projet complet m'a permis de maîtriser l'ensemble du cycle de vie d'une application 
        web moderne, de la conception de la base de données à l'intégration continue, en passant 
        par le développement backend et frontend.
    </p>

    <div class="skills-grid">

        <div class="skill-item">
            <strong>Développement Backend (<a class="link" href="/competences/detail.html?id=7">Spring</a>)</strong>
            <p>
                Mise en place d'une API RESTful robuste avec Spring Boot. Structuration du code en 
                couches (Controller, Service, Repository/Store) pour une meilleure maintenabilité. 
                Gestion de la sécurité et de la configuration de l'application.
            </p>
        </div>

        <div class="skill-item">
            <strong>Développement Frontend (<a class="link" href="/competences/detail.html?id=5">Angular</a>)</strong>
            <p>
                Création d'une application monopage (SPA) dynamique avec <a class="link" href="/competences/detail.html?id=5">Angular</a>. Développement 
                d'une architecture à base de composants (Dashboard, Project, Task), gestion de l'état 
                et communication asynchrone avec le backend via des services.
            </p>
        </div>

        <div class="skill-item">
            <strong>Architecture Full-Stack & Intégration</strong>
            <p>
                Conception de l'API et du contrat de données entre le client et le serveur. 
                Mise en place de la communication (requêtes HTTP), gestion de l'authentification 
                et configuration des règles CORS pour assurer une intégration fluide.
            </p>
        </div>

        <div class="skill-item">
            <strong>Gestion de Base de Données (<a class="link" href="/competences/detail.html?id=6">SQL</a>)</strong>
            <p>
                Modélisation de la structure de la base de données relationnelle (schéma SQL). 
                Création des entités (Projet, Tâche, Utilisateur) et gestion de leurs relations. 
                Utilisation de scripts pour initialiser la base avec des données de test.
            </p>
        </div>

        <div class="skill-item">
            <strong>DevOps & Conteneurisation (<a class="link" href="/competences/detail.html?id=11">Docker</a>)</strong>
            <p>
                Utilisation de <a class="link" href="/competences/detail.html?id=11">Docker</a> et Docker Compose pour créer un environnement de développement 
                isolé et reproductible. Configuration des services (backend, frontend, base de 
                données) pour qu'ils fonctionnent de manière coordonnée.
            </p>
        </div>

        <div class="skill-item">
            <strong>Automatisation & Outillage (Maven, Makefile)</strong>
            <p>
                Utilisation de Maven pour la gestion des dépendances et le build du projet Java. 
                Création d'un Makefile pour simplifier et automatiser les tâches courantes 
                (lancer l'application, exécuter les tests, etc.).
            </p>
            </div>
    </div>
</section>

<!-- =================================================================== -->
<!-- SECTION : CONCLUSION                                                -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Conclusion</b></h3>
    <p>
        PMT est un projet intégrateur qui démontre ma capacité à concevoir, développer et déployer 
        une application web complète et fonctionnelle. Il met en évidence mes compétences techniques 
        en <a class="link" href="/competences/detail.html?id=7">Java/Spring</a> pour la partie serveur et 
        <a class="link" href="/competences/detail.html?id=5">Angular</a> pour la partie client, ainsi 
        que ma maîtrise des outils DevOps modernes comme <a class="link" href="/competences/detail.html?id=11">Docker</a>.  
        Cette réalisation atteste de ma polyvalence et de ma compréhension globale des enjeux liés à 
        la création d'applications robustes, maintenables et prêtes pour la production.
    </p>

    <p>
        Les lendemains du projet s'inscrivent dans une logique d'usage réel : PMT est destiné à 
        devenir mon outil principal pour organiser mon travail et mes projets personnels.  
        J'envisage également de le rendre accessible en ligne, ce qui nécessitera 
        une réarchitecture partielle pour permettre son utilisation par différentes organisations, 
        ainsi qu'un renforcement de la sécurité et de la gestion des accès.  
        Ces évolutions futures permettront de transformer PMT en un outil plus générique, plus 
        sécurisé et potentiellement utilisable par un public plus large.
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
