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
            <b>Java / Spring Boot</b> exposant une API REST, et un frontend en <b>Angular</b> 
            consommant cette API. L'ensemble est conteneurisé avec <b>Docker</b>, garantissant 
            un environnement de développement et de déploiement cohérent et reproductible.
        </p>
    </section>
    
    <!-- =================================================================== -->
    <!-- SECTION : MISSIONS ET COMPÉTENCES DÉVELOPPÉES                       -->
    <!-- =================================================================== -->
    <section class="project-section">
        <h3>Missions et Compétences Développées</h3>
        <p>
            Ce projet complet m'a permis de maîtriser l'ensemble du cycle de vie d'une application 
            web moderne, de la conception de la base de données à l'intégration continue, en passant 
            par le développement backend et frontend.
        </p>
    
        <div class="skills-grid">
    
            <div class="skill-item">
                <strong>Développement Backend (Java & Spring Boot)</strong>
                <p>
                    Mise en place d'une API RESTful robuste avec Spring Boot. Structuration du code en 
                    couches (Controller, Service, Repository/Store) pour une meilleure maintenabilité. 
                    Gestion de la sécurité et de la configuration de l'application.
                </p>
            </div>
    
            <div class="skill-item">
                <strong>Développement Frontend (Angular)</strong>
                <p>
                    Création d'une application monopage (SPA) dynamique avec Angular. Développement 
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
                <strong>Gestion de Base de Données (SQL)</strong>
                <p>
                    Modélisation de la structure de la base de données relationnelle (schéma SQL). 
                    Création des entités (Projet, Tâche, Utilisateur) et gestion de leurs relations. 
                    Utilisation de scripts pour initialiser la base avec des données de test.
                </p>
            </div>
    
            <div class="skill-item">
                <strong>DevOps & Conteneurisation (Docker)</strong>
                <p>
                    Utilisation de Docker et Docker Compose pour créer un environnement de développement 
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
        <h3>Conclusion</h3>
        <p>
            PMT est un projet intégrateur qui démontre ma capacité à concevoir, développer et déployer 
           une application web complète et fonctionnelle. Il met en évidence mes compétences techniques 
           en <b>Java/Spring</b> pour la partie serveur et <b>Angular</b> pour la partie client, ainsi 
           que ma maîtrise des outils DevOps modernes comme <b>Docker</b>.
       </p>
       <p>
           Cette réalisation atteste de ma polyvalence et de ma compréhension globale des enjeux liés à 
           la création d'applications robustes, maintenables et prêtes pour la production.
       </p>
   </section>
<!-- =================================================================== -->
<!-- SECTION : COMPÉTENCES ASSOCIÉES                                     -->
<!-- =================================================================== -->
<section class="project-section">
    <h3>Compétences Associées</h3>
    <div class="competences-list">
        <a href="/competences/detail.html?id=10" class="competence-tag">Spring</a>
        <a href="/competences/detail.html?id=9" class="competence-tag">Java</a>
        <a href="/competences/detail.html?id=7" class="competence-tag">Angular</a>
        <a href="/competences/detail.html?id=2" class="competence-tag">HTML/CSS</a>
        <a href="/competences/detail.html?id=8" class="competence-tag">SQL</a>
        <a href="/competences/detail.html?id=12" class="competence-tag">Docker</a>
        <a href="/competences/detail.html?id=13" class="competence-tag">Kubernetes</a>
        <a href="/competences/detail.html?id=23" class="competence-tag">Esprit critique</a>
        <a href="/competences/detail.html?id=22" class="competence-tag">Créativité</a>
    </div>
</section>
`
