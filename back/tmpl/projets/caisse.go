package tmpl

// CaisseHTML contient la description HTML du projet de refonte de la caisse Escarcelle.
const CaisseHTML = `
<!-- =================================================================== -->
<!-- SECTION : INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <p>
        Ce projet consistait en la refonte complète de l'application de caisse Escarcelle, un outil 
        central utilisé quotidiennement par plus de 600 épiceries solidaires. L'ancienne version, 
        développée en Rebol, était devenue difficile à maintenir, limitée dans ses évolutions et 
        comportait plusieurs risques : absence de chiffrement, interface vieillissante, dépendance 
        à une technologie obsolète.
    </p>

    <p>
        La décision a été prise de repartir de zéro en utilisant 
        <a class="link" href="/competences/detail.html?id=1">Go</a> et la bibliothèque 
        <strong>Fyne</strong>, afin de concevoir une application moderne, multiplateforme 
        (Windows, macOS, Linux), sécurisée et capable de fonctionner aussi bien en ligne qu'en mode 
        autonome. La caisse repose sur une base locale synchronisée avec l'API partenaire, garantissant 
        une continuité de service même en cas de coupure réseau.
    </p>

    <img class="project-image" src="../assets/CaisseEscarcelle.png" alt="Caisse Escarcelle">

    <p>
        Cette refonte a également permis d'optimiser les workflows utilisateurs, 
        d'intégrer de nombreuses fonctionnalités attendues (multi code-barres, remises avancées, 
        tickets par email, mise en attente des ventes, resynchronisation automatique…), et de proposer 
        une interface tactile, plus intuitive, plus accessible et adaptée aux réalités du terrain.
    </p>
</section>

<section class="project-section">
    <h3><b>Contexte et Enjeux</b></h3>
    <p>
        La caisse Escarcelle est un outil critique : elle conditionne le fonctionnement quotidien 
        de centaines d'épiceries solidaires. Toute interruption de service aurait un impact direct 
        sur les bénéficiaires. La refonte devait donc moderniser l'outil sans perturber l'activité, 
        tout en garantissant une transition fluide pour des utilisateurs aux profils variés, dont 
        beaucoup de bénévoles peu familiers avec l'informatique.
    </p>

    <p>
        L'analyse de l'existant a révélé plusieurs limites : absence de chiffrement, interface peu 
        lisible, impossibilité d'évoluer, difficultés de maintenance, et manque de personnalisation. 
        Les retours terrain ont confirmé ces constats : certaines personnes peinaient à lire l'écran, 
        d'autres souhaitaient des remises plus flexibles, un thème personnalisable ou l'envoi des 
        tickets par email.
    </p>

    <p>
        Les enjeux étaient donc multiples : proposer une application <strong>résiliente</strong>, 
        <strong>sécurisée</strong>, <strong>tactile</strong>, <strong>multiplateforme</strong>, 
        capable de fonctionner hors-ligne, parfaitement intégrée à l'écosystème Escarcelle, et 
        déployable sans aucune interruption pour les 625 structures concernées.
    </p>

    <p>
        Ce cadrage a posé des bases solides pour mener une transformation profonde, structurée et 
        maîtrisée, tout en garantissant la continuité de service.
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
            <strong>Moi-même</strong> - en charge de la conception, du développement, 
            des choix techniques, de l'architecture et de la coordination globale du projet.
        </li>

        <li>
            <strong>Brahim Hamdouni</strong> - garant de la cohérence technique et fonctionnelle, 
            apportant un accompagnement régulier, des validations d'architecture et des retours 
            sur les choix de conception.
        </li>

        <li>
            <strong>L'équipe de test</strong> - responsable de la 
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
                Création d'un package d'impression complet : 
                <ul>
                    <li>gestion des tickets ESC/POS pour imprimantes thermiques</li>
                    <li>génération de tickets A4 via appels système</li>
                    <li>configuration avancée (nom/lie d'exemplaires, format, caractères par ligne, test d'impression)</li>
                    <li>intégration avec les paramètres utilisateurs (thème, police, confirmation d'impression)</li>
                </ul>
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
                Implémentation de fonctionnalités essentielles : 
                <ul>
                    <li>saisie client et contrôle des droits d'accès</li>
                    <li>gestion du panier (quantités, remises, suppression, filtrage)</li>
                    <li>mise en attente de ventes multiples</li>
                    <li>encaissement multi-modes de paiement</li>
                    <li>envoi du ticket par email</li>
                    <li>gestion du budget client et des tarifs spécifiques</li>
                    <li>affichage des informations produit (catégorie, stock, DLC, etc.)</li>
                    <li>support complet du multi-code-barres via l'API partenaire</li>
                </ul>
            </p>
        </div>

    </div>
</section>

<!-- =================================================================== -->
<!-- SECTION : CONCLUSION                                                -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Conclusion</b></h3>

    <h5>Résultats</h5>
    <p>
        La refonte de la caisse Escarcelle a été un succès majeur, autant pour l'entreprise que pour 
        moi. Pour Softinnov, elle a permis de moderniser un outil central utilisé par plus de 600 
        épiceries solidaires, d'améliorer la fiabilité du service, de réduire les erreurs en caisse 
        et de fluidifier le travail des équipes sur le terrain. L'application est désormais plus 
        rapide, plus sécurisée, plus ergonomique et beaucoup plus simple à maintenir.
    </p>

    <p>
        Pour ma part, ce projet a été l'un des plus formateurs de mon parcours. Il m'a permis 
        d'acquérir une expertise solide en développement d'applications de bureau en 
        <a class="link" href="/competences/detail.html?id=1">Go</a> avec Fyne, en conception 
        d'architectures événementielles, en gestion de données locales synchronisées, en intégration 
        de périphériques et en pilotage complet d'un projet stratégique, du cadrage à la livraison.
    </p>

    <h5>Pour la suite</h5>
    <p>
        La nouvelle caisse constitue désormais un socle robuste pour l'écosystème Escarcelle. Les 
        évolutions à venir porteront principalement sur des améliorations fonctionnelles, 
        l'enrichissement du paramétrage, l'adaptation aux retours des épiceries et l'optimisation 
        continue de l'expérience utilisateur.  
        Grâce à l'architecture modulaire mise en place, ces évolutions pourront être intégrées de 
        manière fluide, sans remettre en cause la stabilité du système.
    </p>

    <h5>Mon regard critique</h5>
    <p>
        Avec le recul, je considère ce projet comme une réussite solide : l'adoption a été fluide, 
        les retours des utilisateurs très positifs et la transition s'est faite sans interruption de 
        service, ce qui était l'un des enjeux majeurs. La collaboration avec les équipes de test, 
        les épiceries pilotes et le référent technique a été déterminante pour atteindre ce niveau 
        de qualité.
    </p>

    <p>
        Sur le plan technique, j'ai également identifié des axes d'amélioration. La complexité du 
        projet aurait pu être mieux anticipée en amont, notamment en matière de gestion des risques 
        et de découpage des modules. Certaines décisions techniques auraient gagné à être challengées 
        plus tôt, et j'aurais pu documenter davantage certaines parties du code pour faciliter la 
        montée en compétence d'autres développeurs.
    </p>

    <p>
        Malgré ces points d'amélioration, ce projet reste l'un des plus importants de mon parcours. 
        Il m'a permis de progresser autant sur le plan technique que sur le plan organisationnel, et 
        m'a donné une vision plus mature de ce que représente la refonte d'un outil métier critique 
        à grande échelle.
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
