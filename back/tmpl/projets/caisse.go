package tmpl

// CaisseHTML contient la description HTML du projet de refonte de la caisse Escarcelle.
const CaisseHTML = `
<!-- =================================================================== -->
<!-- SECTION : INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <p>
        Ce projet consistait en la refonte complète de l'application de caisse de l'épicerie solidaire Escarcelle. 
        L'ancienne version, développée en Rebol, était devenue obsolète : elle ne supportait pas les protocoles 
        sécurisés (SSL/TLS), son interface était vieillissante, et sa maintenance devenait de plus en plus complexe.
    </p>

    <p>
        La décision a été prise de repartir de zéro en utilisant <a class="link" href="/competences/detail.html?id=1">Go</a> et la bibliothèque <strong>Fyne</strong> afin de créer 
        une application de bureau moderne, multiplateforme (Windows, macOS, Linux), sécurisée et capable de 
        fonctionner en mode connecté comme en mode autonome. La caisse repose sur une base de données locale 
        synchronisée avec l'API partenaire, permettant une continuité de service même en cas de coupure Internet.
    </p>

    <p>
        Cette refonte a également été l'occasion de repenser entièrement les workflows utilisateurs, 
        d'intégrer de nouvelles fonctionnalités (multi-code-barres, gestion avancée des tickets, 
        resynchronisation automatique, mise en attente des ventes, etc.) et de proposer une interface 
        adaptée aux écrans tactiles, plus intuitive et plus robuste.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : ACTEURS ET INTERACTIONS DU PROJET                        -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Acteurs et Interactions</b></h3>
    <p>
        Le projet a mobilisé plusieurs acteurs aux rôles complémentaires :
    </p>

    <ul>
        <li>
            <strong>Moi-même (développeur principal)</strong> - en charge de la conception, du développement, 
            des choix techniques, de l'architecture et de la coordination globale du projet.
        </li>

        <li>
            <strong>Mon maître d'apprentissage</strong> - garant de la cohérence technique et fonctionnelle, 
            apportant un accompagnement régulier, des validations d'architecture et des retours 
            sur les choix de conception.
        </li>

        <li>
            <strong>Une testeuse dédiée</strong> - responsable de la qualification fonctionnelle, de la 
            détection des anomalies, de la validation des workflows utilisateurs et de la 
            conformité aux besoins métier.
        </li>

        <li>
            <strong>Les épiceries solidaires (utilisateurs finaux)</strong> - parties prenantes essentielles 
            pour la compréhension des usages réels, la remontée des irritants, la priorisation 
            des fonctionnalités et la validation terrain.
        </li>
    </ul>

    <p>
        Les interactions entre ces acteurs ont rythmé le projet : ateliers de cadrage, 
        démonstrations intermédiaires, retours utilisateurs, sessions de tests, et ajustements 
        continus pour garantir une solution robuste, intuitive et adaptée aux contraintes du terrain.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISSIONS ET COMPÉTENCES DÉVELOPPÉES                       -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Missions et Compétences Développées</b></h3>
    <p>
        L'application repose sur une architecture modulaire et événementielle permettant de découpler 
        l'interface, la logique métier, la synchronisation réseau et la gestion des données locales. 
        Mes missions ont couvert l'ensemble du cycle de développement, de la conception à la livraison, 
        en passant par la définition des protocoles d'échange avec l'API partenaire.
    </p>

    <div class="skills-grid">

        <div class="skill-item">
            <strong>Interface Graphique (Fyne)</strong>
            <p>
                Conception et implémentation d'une interface utilisateur complète, tactile-friendly, 
                intégrant clavier virtuel, navigation par écrans, gestion des états, et affichage dynamique 
                des informations client, panier et produits.  
                Développement de composants personnalisés pour la saisie des codes-barres, la gestion 
                des remises, la mise en attente des ventes et l'encaissement multi-modes de paiement.
            </p>
        </div>

        <div class="skill-item">
            <strong>Gestion de Données (<a class="link" href="/competences/detail.html?id=6">SQLite</a>)</strong>
            <p>
                Mise en place du package <i>store</i> pour gérer la base locale SQLite : structure des tables, 
                migrations, indexation, gestion des ventes en attente, stockage des tickets, et 
                synchronisation bidirectionnelle avec l'API.  
                Conception d'un mécanisme de résilience permettant de conserver les ventes en cas de 
                coupure réseau et de les renvoyer automatiquement à la reconnexion.
            </p>
        </div>

        <div class="skill-item">
            <strong>Connectivité et Réseau</strong>
            <p>
                Développement du module <i>network</i> pour gérer la communication avec l'API partenaire : 
                authentification, gestion des erreurs, timeouts configurables, support du proxy, 
                synchronisation des données (clients, produits, tarifs, stocks) et envoi des ventes en temps réel.  
                Mise en place d'un indicateur de connexion en direct et d'un système de reprise automatique 
                des synchronisations échouées.
            </p>
        </div>

        <div class="skill-item">
            <strong>Intégration d'Impression</strong>
            <p>
                Création d'un package d'impression complet : <br>
                • gestion des tickets ESC/POS pour imprimantes thermiques <br>
                • génération de tickets A4 via appels système <br>
                • configuration avancée (nombre d'exemplaires, format, caractères par ligne, test d'impression) <br>
                • intégration avec les paramètres utilisateurs (thème, police, confirmation d'impression) <br>
            </p>
        </div>

        <div class="skill-item">
            <strong>Architecture Événementielle</strong>
            <p>
                Conception du cœur de l'application autour d'un bus d'événements permettant de gérer 
                les transitions d'écrans, les interactions utilisateur, les mises à jour du panier, 
                les changements d'état réseau et les notifications sonores.  
                Cette architecture a permis une séparation claire des responsabilités et une grande 
                facilité d'évolution.
            </p>
        </div>

        <div class="skill-item">
            <strong>Périphériques Audio</strong>
            <p>
                Intégration d'un module audio pour les notifications (erreurs, confirmations, alertes), 
                configurable par l'utilisateur et compatible avec les différentes plateformes.
            </p>
        </div>

        <div class="skill-item">
            <strong>Fonctionnalités Métier Avancées</strong>
            <p>
                Implémentation de fonctionnalités essentielles : <br> 
                • saisie client et contrôle des droits d'accès <br>
                • gestion du panier (quantités, remises, suppression, filtrage) <br>
                • mise en attente de ventes multiples <br>
                • encaissement multi-modes de paiement <br>
                • envoi du ticket par email <br>
                • gestion du budget client et des tarifs spécifiques <br>
                • affichage des informations produit (catégorie, stock, DLC, etc.) <br>
                • support complet du multi-code-barres via l'API partenaire <br>
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
        Cette refonte a été un succès technique majeur, aboutissant à une application de caisse moderne, 
        rapide, sécurisée et considérablement plus simple à maintenir et à faire évoluer.  
        Le projet m'a permis d'acquérir une expertise approfondie dans le développement d'applications 
        de bureau en <a class="link" href="/competences/detail.html?id=1">Go</a> avec Fyne, la conception d'architectures événementielles, la gestion de données 
        locales synchronisées et l'intégration de périphériques (imprimantes, douchettes, audio).  
        Il constitue aujourd'hui un socle robuste pour les futures évolutions de l'écosystème Escarcelle.
    </p>

    <p> 
        À l'issue du projet, l'application constitue désormais un socle robuste pour l'écosystème Escarcelle. 
        Les lendemains du projet se concentrent principalement sur la maintenance évolutive et corrective, 
        l'ajout de petites améliorations fonctionnelles, ainsi que l'adaptation continue aux retours des épiceries. 
        La structure modulaire mise en place facilite grandement ces évolutions futures. 
    </p>
</section>
<!-- =================================================================== -->
<!-- SECTION : COMPÉTENCES ASSOCIÉES                                     -->
<!-- =================================================================== -->
<section class="project-section">
    <h5>Compétences Associées</h5>
    <div class="competences-list">
        <a href="/competences/detail.html?id=1" class="competence-tag technique">Go</a>
        <a href="/competences/detail.html?id=6" class="competence-tag technique">SQL</a>
        <a href="/competences/detail.html?id=11" class="competence-tag devops">Docker</a>
        <a href="/competences/detail.html?id=13" class="competence-tag humain">Gestion de projet</a>
        <a href="/competences/detail.html?id=15" class="competence-tag humain">Créativité</a>
        <a href="/competences/detail.html?id=14" class="competence-tag humain">Flexibilité</a>
        <a href="/competences/detail.html?id=16" class="competence-tag humain">Esprit critique</a>
    </div>
</section>
`
