package tmpl

// CaisseProject contient la description HTML du projet de refonte de la caisse Escarcelle.
const CaisseProject = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <p>
        Ce projet consistait en la refonte complète de l'application de caisse de l'épicerie solidaire Escarcelle. L'ancienne version, développée en Rebol, était devenue obsolète : elle ne supportait pas les protocoles sécurisés (SSL/TLS), son interface était vieillissante et sa maintenance complexe.
    </p>
    <p>
        La décision a été prise de repartir de zéro en utilisant <b>Go</b> et la bibliothèque <b>Fyne</b> pour créer une application de bureau moderne, multiplateforme (Windows, macOS, Linux), sécurisée et capable de fonctionner avec ou sans connexion internet (online/offline) grâce à une base de données locale synchronisée.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISSIONS ET COMPÉTENCES DÉVELOPPÉES                       -->
<!-- =================================================================== -->
<section class="project-section">
    <h3>Missions et Compétences Développées</h3>
    <p>
        L'architecture de l'application repose sur un système à base d'événements permettant de découpler les différents modules et de gérer les états de l'interface. Mes missions ont couvert le cycle de vie complet du développement :
    </p>
    <div class="skills-grid">
        <div class="skill-item">
            <strong>Interface Graphique (Fyne)</strong>
            <p>Conception et implémentation de toute l'interface utilisateur multiplateforme avec la bibliothèque Fyne en Go.</p>
        </div>
        <div class="skill-item">
            <strong>Gestion de Données (SQLite)</strong>
            <p>Mise en place du package "store" pour gérer la base de données locale SQLite et assurer la synchronisation avec l'API partenaire.</p>
        </div>
        <div class="skill-item">
            <strong>Connectivité et Réseau</strong>
            <p>Développement du module "network" pour gérer la communication avec l'API, incluant la gestion de proxy.</p>
        </div>
        <div class="skill-item">
            <strong>Intégration d'Impression</strong>
            <p>Création d'un package d'impression gérant les tickets de caisse (commandes ESC/POS) et les factures A4 (via syscalls).</p>
        </div>
        <div class="skill-item">
            <strong>Architecture Événementielle</strong>
            <p>Conception du cœur de l'application basé sur les événements pour gérer les changements d'écrans et les états de la caisse.</p>
        </div>
        <div class="skill-item">
            <strong>Périphériques Audio</strong>
            <p>Intégration d'un module pour émettre des signaux sonores informant l'utilisateur.</p>
        </div>
    </div>
</section>

<!-- =================================================================== -->
<!-- SECTION : CONCLUSION                                                -->
<!-- =================================================================== -->
<section class="project-section">
    <h3>Conclusion</h3>
    <p>
        Cette refonte a été un succès technique, aboutissant à une application de caisse rapide, sécurisée et beaucoup plus facile à maintenir et à faire évoluer. Ce projet m'a permis de développer une expertise solide dans la création d'applications de bureau en Go avec Fyne et dans la conception d'architectures logicielles résilientes.
    </p>
</section>
`
