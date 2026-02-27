package tmpl

// EscarcelleHTML contient la description HTML du projet Escarcelle.
const EscarcelleHTML = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <p>
        <b>Escarcelle</b> est une solution logicielle complète dédiée à la gestion des épiceries solidaires. 
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
<!-- SECTION : ACTEURS ET INTERACTIONS DU PROJET                        -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Acteurs et Interactions</b></h3>
    <p>
        Le projet Escarcelle a impliqué plusieurs acteurs aux rôles complémentaires, chacun contribuant 
        à la modernisation progressive de l'écosystème :
    </p>

    <ul>
        <li>
            <strong>Mon maître d'apprentissage (développeur principal)</strong> - responsable de la vision technique 
            globale, des choix d'architecture, de la priorisation des évolutions et de la coordination 
            entre les différents modules.
        </li>

        <li>
            <strong>Moi-même (développeur full-stack)</strong> - en charge de la modernisation du socle existant, 
            de la refonte de modules critiques, de la création de nouveaux services et de la migration 
            progressive du legacy vers la stack moderne <a class="link" href="/competences/detail.html?id=1">Go</a>/<a class="link" href="/competences/detail.html?id=4">Svelte</a>.
        </li>

        <li>
            <strong>Une testeuse dédiée</strong> - intervenant sur la qualification fonctionnelle, la validation 
            des workflows, la détection des anomalies et la conformité aux besoins métier.
        </li>

        <li>
            <strong>Les épiceries solidaires (utilisateurs finaux)</strong> - apportant des retours terrain essentiels 
            pour orienter les priorités, identifier les irritants et valider les évolutions.
        </li>
    </ul>

    <p>
        Les interactions entre ces acteurs ont rythmé le projet : ateliers de cadrage, tests utilisateurs, 
        démonstrations régulières, retours terrain et ajustements continus. Cette collaboration a permis 
        d'aligner les choix techniques avec les usages réels et les contraintes opérationnelles des 
        épiceries solidaires.
    </p>
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
            <strong>Base de Données (<a class="link" href="/competences/detail.html?id=6">SQL</a>)</strong>
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
            <strong>Legacy (<a class="link" href="/competences/detail.html?id=9">Rebol</a>)</strong>
            <p>
                • Mise à jour du code <a class="link" href="/competences/detail.html?id=9">Rebol</a> de l'application de gestion et de l'ancienne caisse afin 
                d'intégrer la nouvelle structure de catégories.<br>
                • Maintenance corrective et évolutive sur un codebase historique nécessitant une 
                compréhension fine du fonctionnement interne et des contraintes métier.<br>
                • Décommission progressive de certaines parties du legacy grâce à la création de nouveaux 
                composants modernes (ex : gestion des codes-barres).
            </p>
        </div>

        <div class="skill-item">
            <strong>Backend (<a class="link" href="/competences/detail.html?id=1">Go</a>)</strong>
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
            <strong>Frontend (<a class="link" href="/competences/detail.html?id=4">Svelte</a>)</strong>
            <p>
                • Développement des interfaces de l'application web de gestion, avec un accent particulier 
                sur l'ergonomie, la clarté des workflows et la cohérence de l'expérience utilisateur.<br>
                • Refonte complète de la gestion des produits pour intégrer le multi-code-barres via des 
                composants <a class="link" href="/competences/detail.html?id=4">Svelte</a> modernes, remplaçant le code 
                <a class="link" href="/competences/detail.html?id=9">Rebol</a> historique.<br>
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

    <p>
        Les lendemains du projet s'inscrivent dans une dynamique d'amélioration continue : accompagnement 
        des épiceries dans leur utilisation quotidienne, ajout d'évolutions pertinentes et poursuite de 
        la migration du code <a class="link" href="/competences/detail.html?id=9">Rebol</a> vers des composants modernes. 
        Un module majeur est actuellement en phase de validation : l'intégration d'OpenFoodFacts, 
        permettant aux épiceries de référencer automatiquement des produits à partir du code-barres, 
        de manière standardisée et fiable. Ce travail, combiné à la réduction progressive du legacy, 
        prépare Escarcelle à devenir une plateforme encore plus robuste, interopérable et durable.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : COMPÉTENCES ASSOCIÉES                                     -->
<!-- =================================================================== -->
<section class="project-section">
    <h5>Compétences Associées</h5>
    <div class="competences-list">
        <a href="/competences/detail.html?id=1" class="competence-tag technique">Go</a>
        <a href="/competences/detail.html?id=9" class="competence-tag technique">Rebol</a>
        <a href="/competences/detail.html?id=2" class="competence-tag technique">HTML/CSS</a>
        <a href="/competences/detail.html?id=4" class="competence-tag technique">Svelte</a>
        <a href="/competences/detail.html?id=6" class="competence-tag technique">SQL</a>
        <a href="/competences/detail.html?id=12" class="competence-tag devops">Bash</a>
        <a href="/competences/detail.html?id=11" class="competence-tag devops">Docker</a>
        <a href="/competences/detail.html?id=13" class="competence-tag humain">Gestion de projet</a>
        <a href="/competences/detail.html?id=14" class="competence-tag humain">Flexibilité</a>
    </div>
</section>
`
