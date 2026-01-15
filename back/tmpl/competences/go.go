package tmpl

// GoComp contient la description HTML de la compétence Go.
const GoComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        Go, souvent appelé Golang, est un langage de programmation compilé et concurrentiel développé par Google. 
        Il est reconnu pour sa simplicité, sa performance et sa robustesse, ce qui en fait un choix privilégié 
        pour le développement de services backend, d'outils en ligne de commande et d'infrastructures cloud-native.
    </p>
    <p>
        Mon expérience avec Go couvre l'ensemble du cycle de développement d'applications, de la conception 
        d'API RESTful à la création d'applications de bureau complexes. J'apprécie particulièrement sa 
        gestion native de la concurrence, son typage statique qui garantit la sécurité du code, et son 
        écosystème d'outils très complet.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Mon expertise en Go s'est concrétisée à travers des projets d'envergure où j'ai pu mettre en œuvre un large éventail de fonctionnalités, notamment dans le cadre des projets <strong>Escarcelle</strong> et <strong>Caisse Escarcelle</strong> :
    </p>
    <ul>
        <li>
            <strong>Développement d'applications de bureau :</strong> Utilisation du toolkit UI multiplateforme Fyne pour construire l'intégralité de l'interface utilisateur de l'application de caisse, incluant la gestion des layouts, des événements, et le développement de widgets personnalisés.
        </li>
        <li>
            <strong>Création d'APIs et de logiques métier :</strong> Conception et développement d'APIs RESTful robustes pour servir de backend, et implémentation de logiques métier complexes spécifiques au secteur de la distribution (gestion des stocks, des promotions, des tarifs, etc.).
        </li>
        <li>
            <strong>Intégration de périphériques et impression :</strong> Développement de modules d'impression complets pour la génération de documents A4 (via des appels système) et de tickets de caisse sur imprimantes thermiques avec le protocole <strong>ESC/POS</strong>.
        </li>
        <li>
            <strong>Qualité et fiabilité du code :</strong> Mise en place de stratégies de tests complètes, incluant des <strong>tests unitaires</strong> pour valider chaque fonction métier et des <strong>tests end-to-end</strong> pour garantir la stabilité et la robustesse des applications en conditions réelles.
        </li>
    </ul>
</section>
<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=2">Caisse</a></li>
        </ul>
    </div>
</section>
`
