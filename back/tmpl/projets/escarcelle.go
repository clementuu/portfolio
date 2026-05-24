package tmpl

// EscarcelleHTML contient la description HTML du projet Escarcelle.
const EscarcelleHTML = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <h3>Présentation</h3>

    <p>
        <strong>Escarcelle</strong> est un système d'information complet dédié aux épiceries solidaires. 
        Initialement conçu comme un outil de gestion interne, il a progressivement évolué vers un 
        véritable écosystème numérique regroupant une application web moderne, une caisse connectée 
        et une API ouverte aux partenaires externes.
    </p>

    <p>
        Mon rôle a été d'accompagner cette modernisation en refondant des modules historiques, en 
        introduisant une stack plus robuste (<a class="link" href="/competences/detail.html?id=1">Go</a> / 
        <a class="link" href="/competences/detail.html?id=4">Svelte</a>) et en améliorant la 
        maintenabilité, la performance et l'expérience utilisateur. Ce travail a impliqué des 
        migrations sensibles, des refontes structurelles et la création de nouveaux services 
        interopérables.
    </p>

    <img class="project-image" src="../assets/AccueilEscarcelle.png" alt="Accueil Escarcelle">

    <p>
        L'écosystème Escarcelle repose sur trois modules principaux :
    </p>

    <ul class="escarcelle-list">
        <li>
            <strong>Application de Gestion</strong>
            Gestion des stocks, bénéficiaires, produits, rapports.
        </li>
        <li>
            <strong>Logiciel de Caisse</strong> 
            Enregistrement des ventes en boutique, synchronisée avec l'API.
        </li>
        <li>
            <strong>API Partenaire</strong>
            Interface sécurisée pour les intégrations externes (ex : click & collect).
        </li>
    </ul>
</section>

<section class="project-section">
    <h3>Contexte et Enjeux</h3>
    <p>
        Escarcelle est utilisé quotidiennement par plus de 625 épiceries solidaires, représentant 
        plus de 1,42 million de distributions et plus de 114 000 familles bénéficiaires. Moderniser 
        un tel système impliquait de préserver la continuité de service tout en améliorant un 
        socle historique devenu difficile à maintenir.
    </p>

    <p>
        Le principal enjeu était de faire cohabiter un code legacy en 
        <a class="link" href="/competences/detail.html?id=9">Rebol</a> avec une stack moderne 
        <a class="link" href="/competences/detail.html?id=1">Go</a>/<a class="link" href="/competences/detail.html?id=4">Svelte</a>, 
        tout en garantissant la compatibilité des différents modules et la stabilité du système. Les épiceries 
        attendaient également des évolutions concrètes : meilleure gestion des produits, workflows 
        simplifiés, réduction des erreurs et accélération des opérations.
    </p>

    <p>
        Les enjeux techniques étaient tout aussi importants : migration de MySQL 5.5 vers 8.0, 
        refonte de modules critiques, introduction du multi code-barres, intégration de la base de données 
        d'OpenFoodFacts, restructuration de l'architecture et création de nouveaux outils CLI pour améliorer 
        l'environnement de développement.
    </p>

    <p>
        L'objectif global était clair : rendre Escarcelle plus robuste, plus moderne, plus 
        interopérable et plus simple à utiliser, tout en renforçant son impact social auprès des 
        bénéficiaires.
    </p>

    <h3>Acteurs et Interactions</h3>
    <p>
        Le projet Escarcelle a impliqué plusieurs acteurs aux rôles complémentaires, chacun contribuant 
        à la modernisation progressive de l'écosystème :
    </p>

    <ul class="actors-list">
        <li>
            <strong>Brahim Hamdouni</strong>
            Responsable de la vision technique globale, des choix d'architecture, de la priorisation 
            des évolutions et de la coordination entre les différents modules.
        </li>

        <li>
            <strong>Moi-même</strong>
            En charge de la modernisation du socle existant, de la refonte de modules critiques, de la 
            création de nouveaux services et de la migration progressive du legacy vers la stack moderne 
            <a class="link" href="/competences/detail.html?id=1">Go</a>/<a class="link" href="/competences/detail.html?id=4">Svelte</a>.
        </li>

        <li>
            <strong>L'équipe de test</strong>
            Intervenant sur la qualification fonctionnelle, la détection des anomalies et la conformité 
            aux besoins métier.
        </li>

        <li>
            <strong>Les épiceries solidaires</strong>
            Apportant des retours terrain essentiels pour orienter les priorités, identifier les irritants 
            et valider les évolutions.
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
    <h3>Missions et Compétences Développées</h3>
    <p>
        Mon intervention sur Escarcelle a couvert un large spectre technique, allant de la modernisation 
        d'un socle existant à la conception de nouveaux modules. Ce projet m'a permis de renforcer mes 
        compétences en architecture logicielle, en migration de systèmes et en développement full-stack.
    </p>

    <div class="skills-grid">
        <div class="skill-item">
            <strong>Base de Données (<a class="link" href="/competences/detail.html?id=6">SQL</a>)</strong>
            <p>
                <ul>
                    <li>
                        Migration complète de MySQL 5 vers MySQL 8, incluant l'analyse d'impact, 
                        la mise à jour des requêtes incompatibles et l'adaptation du code applicatif.
                    </li>
                    <li>
                        Refonte du système de catégories produits : passage d'un modèle basé sur l'ordre des IDs 
                        à une structure hiérarchique moderne utilisant un <i>parent_id</i>, rendant l'outil plus 
                        flexible et maintenable.
                    </li>
                    <li>
                        Conception du système multi-code-barres : création d'une table dédiée permettant 
                    d'associer un nombre illimité de codes à un produit, remplaçant l'ancienne limite de deux 
                    codes par article.
                    </li>
            </p>
        </div>

        <div class="skill-item">
            <strong>Legacy (<a class="link" href="/competences/detail.html?id=9">Rebol</a>)</strong>
            <p>
                <ul>
                    <li>
                        Mise à jour du code <a class="link" href="/competences/detail.html?id=9">Rebol</a> de l'application de gestion et de l'ancienne caisse afin 
                        d'intégrer la nouvelle structure de catégories.
                    </li>
                    <li>
                        Maintenance corrective et évolutive sur un codebase historique nécessitant une 
                        compréhension fine du fonctionnement interne et des contraintes métier.
                    </li>
                    <li>
                        Décommission progressive de certaines parties du legacy grâce à la création de nouveaux 
                        composants modernes (ex : gestion des codes-barres).
                    </li>
                </ul>
            </p>
        </div>

        <div class="skill-item">
            <strong>Backend (<a class="link" href="/competences/detail.html?id=1">Go</a>)</strong>
            <p>
                <ul>
                    <li>
                        Conception et développement d'une API partenaire RESTful permettant à une entreprise 
                        tierce de créer une application de click & collect.
                    </li>
                    <li>
                        Extension de l'API pour supporter le système multi-code-barres, incluant la gestion 
                        CRUD des codes associés à un produit.
                    </li>
                    <li>
                        Mise en place d'une architecture sécurisée, documentée et extensible, destinée également 
                        à alimenter la future application de caisse.
                    </li>
                </ul>
            </p>
        </div>

        <div class="skill-item">
            <strong>Frontend (<a class="link" href="/competences/detail.html?id=4">Svelte</a>)</strong>
            <p>
                <ul>
                    <li>
                        Développement des interfaces de l'application web de gestion, avec un accent particulier 
                        sur l'ergonomie, la clarté des workflows et la cohérence de l'expérience utilisateur.
                    </li>
                    <li>
                        Refonte complète de la gestion des produits pour intégrer le multi-code-barres via des 
                        composants <a class="link" href="/competences/detail.html?id=4">Svelte</a> modernes, remplaçant le code 
                        <a class="link" href="/competences/detail.html?id=9">Rebol</a> historique.
                    </li>
                    <li>
                        Amélioration de la maintenabilité et de la lisibilité du code grâce à une architecture 
                        front plus modulaire.
                    </li>
                </ul>
            </p>
        </div>

        <div class="skill-item">
            <strong>Architecture Logicielle</strong>
            <p>
                <ul>
                    <li>
                        Réorganisation complète de l'architecture du projet en recentrant les entités et les règles métier à la racine du codebase, afin de les placer au cœur de l'écosystème applicatif.
                    </li>
                    <li>
                        Mutualisation des règles métier entre les différents modules (gestion, caisse, API), supprimant la duplication de logique et améliorant la cohérence globale.
                    </li>
                    <li>
                        Analyse approfondie des risques liés à cette refonte (impacts sur le legacy, compatibilité ascendante, dépendances croisées) et conduite d'une refactorisation majeure de plusieurs modules pour garantir une transition progressive et sécurisée.
                    </li>
                </ul>
            </p>
        </div>

        <div class="skill-item">
            <strong>Intégration Open Food Facts</strong>
            <p>
                Développement d'un outil complet permettant d'importer et transformer les données 
                d'OpenFoodFacts pour accélérer le référencement des produits.  
                Cet outil inclus un algorithme de scoring qui proposer automatiquement la catégorie 
                Escarcelle la plus pertinente.  
            </p>
            <p>
                Intégration dans l'application : préremplissage automatique des fiches produits à 
                partir du code-barres.
            </p>
        </div>

        <div class="skill-item">
            <strong>Outils CLI & Environnement de Développement</strong>
            <p>
                Création d'applications en ligne de commande pour automatiser des tâches complexes 
                (import OFF, vérifications d'intégrité des données, migrations).  
                Amélioration de l'environnement de développement : scripts, automatisations, 
                outils internes facilitant les tests, la maintenance et la montée en compétence 
                des nouveaux développeurs.
            </p>
        </div>
    </div>
</section>

<!-- =================================================================== -->
<!-- SECTION : CONCLUSION                                                -->
<!-- =================================================================== -->
<section class="project-section">
    <h3>Résultats</h3>
    <p>
        Escarcelle est aujourd'hui une plateforme plus moderne, plus performante et plus simple à 
        maintenir. Les épiceries bénéficient d'outils plus fiables, d'une gestion des produits 
        accélérée (multi-code-barres, OpenFoodFacts), et d'une expérience utilisateur plus fluide.  
        Pour Softinnov, cette modernisation renforce la pérennité du système et ouvre la voie à de 
        nouvelles intégrations externes.
    </p>

    <p>
        Pour moi, ce projet a été un terrain d'apprentissage exceptionnel : migrations sensibles, 
        architecture logicielle, développement full-stack, gestion du legacy, coordination technique 
        et conception de modules à fort impact métier.
    </p>

    <h3>Pour la suite</h3>
    <p>
        Les évolutions à venir concernent principalement la poursuite de la migration du legacy, 
        l'amélioration continue des workflows avec la refonte de certains modules et processus, 
        l'enrichissement du paramétrage et l'intégration progressive de nouveaux services partenaires. 
        L'architecture modernisée permet désormais d'ajouter des fonctionnalités de manière plus 
        fluide et plus sécurisée.
    </p>

    <h3>Mon regard critique</h3>
    <p>
        Je considère ce projet comme une réussite solide, autant sur le plan technique que sur 
        l'impact social. La nouvelle architecture, les migrations et les différentes évolutions 
        ont permis de sécuriser l'avenir de la plateforme, tout en facilitant le quotidien des équipes.
    </p>

    <p>
        Il y a toutefois quelques points à améliorer : certaines migrations auraient pu être 
        mieux découpées, une documentation plus détaillée aurait été bienvenue 
        pour faciliter la transition entre legacy et stack moderne. 
        Ce projet reste tout de même l'un des plus formateurs de mon parcours. Il m'a permis 
        d'appréhender différents aspects du développement et de la modernisation de solutions informatiques, en 
        travaillant sur un outil utilisé chaque jour par des milliers de personnes.
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
