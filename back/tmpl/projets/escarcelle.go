package tmpl

// EscarcelleHTML contient la description HTML du projet Escarcelle.
const EscarcelleHTML = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <p>
        Escarcelle est une solution logicielle complète dédiée à la gestion des épiceries solidaires. 
        Conçu initialement comme un outil de gestion interne, le projet a progressivement évolué vers 
        un véritable écosystème numérique regroupant une application web moderne, une caisse connectée 
        et une API ouverte aux partenaires externes.
    </p>

    <p>
        Mon rôle a été d'accompagner cette transformation en modernisant un socle historique, en 
        introduisant des technologies plus robustes et en repensant plusieurs modules clés pour 
        améliorer la maintenabilité, la performance et l'expérience utilisateur. Ce travail a impliqué 
        des migrations sensibles, des refontes structurelles et la création de nouveaux services 
        interopérables.
    </p>

    <p>
        La plateforme se compose de trois modules principaux :
    </p>

    <ul>
        <li><b>Application Web de Gestion :</b> Interface centrale permettant la gestion des stocks, des bénéficiaires, des produits et la génération de rapports.</li>
        <li><b>Application de Caisse :</b> Logiciel de point de vente utilisé en boutique pour enregistrer les transactions.</li>
        <li><b>API Partenaire :</b> API sécurisée destinée aux intégrations externes, notamment pour le développement d'applications tierces (ex : click & collect).</li>
    </ul>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISSIONS ET COMPÉTENCES DÉVELOPPÉES                       -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Missions et Compétences Développées</b></h3>
    <p>
        Mon intervention sur Escarcelle a couvert un large spectre technique, allant de la modernisation 
        d'un socle existant à la conception de nouveaux modules. Ce projet m'a permis de renforcer mes 
        compétences en architecture logicielle, en migration de systèmes et en développement full-stack.
    </p>

    <div class="skills-grid">

        <div class="skill-item">
            <strong>Base de Données (SQL)</strong>
            <p>
                • Migration complète de MySQL 5 vers MySQL 8, incluant l'analyse d'impact, la mise à jour 
                des requêtes incompatibles et l'adaptation du code applicatif.<br>
                • Refonte du système de catégories produits : passage d'un modèle basé sur l'ordre des IDs 
                à une structure hiérarchique moderne utilisant un <i>parent_id</i>, rendant l'outil plus 
                flexible et maintenable.<br>
                • Conception du système multi-code-barres : création d'une table dédiée permettant 
                d'associer un nombre illimité de codes à un produit, remplaçant l'ancienne limite de deux 
                codes par article.
            </p>
        </div>

        <div class="skill-item">
            <strong>Legacy (Rebol)</strong>
            <p>
                • Mise à jour du code Rebol de l'application de gestion et de l'ancienne caisse afin 
                d'intégrer la nouvelle structure de catégories.<br>
                • Maintenance corrective et évolutive sur un codebase historique nécessitant une 
                compréhension fine du fonctionnement interne et des contraintes métier.<br>
                • Décommission progressive de certaines parties du legacy grâce à la création de nouveaux 
                composants modernes (ex : gestion des codes-barres).
            </p>
        </div>

        <div class="skill-item">
            <strong>Backend (Go)</strong>
            <p>
                • Conception et développement d'une API partenaire RESTful permettant à une entreprise 
                tierce de créer une application de click & collect.<br>
                • Extension de l'API pour supporter le système multi-code-barres, incluant la gestion 
                CRUD des codes associés à un produit.<br>
                • Mise en place d'une architecture sécurisée, documentée et extensible, destinée également 
                à alimenter la future application de caisse.
            </p>
        </div>

        <div class="skill-item">
            <strong>Frontend (Svelte)</strong>
            <p>
                • Développement des interfaces de l'application web de gestion, avec un accent particulier 
                sur l'ergonomie, la clarté des workflows et la cohérence de l'expérience utilisateur.<br>
                • Refonte complète de la gestion des produits pour intégrer le multi-code-barres via des 
                composants Svelte modernes, remplaçant le code Rebol historique.<br>
                • Amélioration de la maintenabilité et de la lisibilité du code grâce à une architecture 
                front plus modulaire.
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
        Le projet Escarcelle a été une expérience particulièrement enrichissante, mêlant modernisation 
        d'un système existant, création de nouveaux outils et collaboration avec des acteurs externes. 
        Il m'a permis de travailler sur des problématiques concrètes à fort impact social, tout en 
        consolidant mes compétences en architecture logicielle, en migration de données et en 
        développement full-stack.  
        Cette expérience m'a également appris à faire cohabiter du legacy et des technologies modernes 
        au sein d'un écosystème cohérent, évolutif et durable.
    </p>
</section>
<!-- =================================================================== -->
<!-- SECTION : COMPÉTENCES ASSOCIÉES                                     -->
<!-- =================================================================== -->
<section class="project-section">
    <h5>Compétences Associées</h5>
    <div class="competences-list">
        <a href="/competences/detail.html?id=1" class="competence-tag technique">Go</a>
        <a href="/competences/detail.html?id=10" class="competence-tag technique">Rebol</a>
        <a href="/competences/detail.html?id=2" class="competence-tag technique">HTML/CSS</a>
        <a href="/competences/detail.html?id=4" class="competence-tag technique">Svelte</a>
        <a href="/competences/detail.html?id=6" class="competence-tag technique">SQL</a>
        <a href="/competences/detail.html?id=13" class="competence-tag devops">Bash</a>
        <a href="/competences/detail.html?id=12" class="competence-tag devops">Docker</a>
        <a href="/competences/detail.html?id=14" class="competence-tag humain">Gestion de projet</a>
        <a href="/competences/detail.html?id=15" class="competence-tag humain">Flexibilité</a>
    </div>
</section>
`
