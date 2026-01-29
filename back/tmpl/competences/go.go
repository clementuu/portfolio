package tmpl

// GoCompetence contient la description HTML de la compétence Go.
const GoCompetence = `
<section class="comp-section">
    <h3><b>Définition</b></h3>
    <p>
        Go, souvent appelé Golang, est un langage de programmation compilé et concurrentiel développé par Google. 
        Il est reconnu pour sa simplicité, sa performance et sa robustesse, ce qui en fait un choix privilégié 
        pour le développement de services backend, d'outils en ligne de commande et d'infrastructures cloud-native.
    </p>
    <p>
        Mon expérience avec Go couvre l'ensemble du cycle de développement d'applications, de la conception 
        d'API RESTful à la création d'applications de bureau complexes. J'apprécie particulièrement sa 
        gestion native de la concurrence, son typage statique qui garantit la sécurité du code, et son 
        écosystème d'outils très complet. Son approche minimaliste et sa capacité à produire des binaires 
        autonomes en font un langage très pertinent dans l'écosystème des microservices et du cloud.
    </p>
</section>

<section class="comp-section">
    <h3><b>Éléments de preuve</b></h3>
    <p>
        Dans le cadre des projets Escarcelle et Caisse, 
        j'ai développé l'intégralité du backend et de l'application de caisse en Go. 
        Cette expérience a prouvé ma maîtrise de Go pour des applications critiques.
    </p>
    
    <p>
        <strong>Développement d'une application de bureau multiplateforme (<a class="link" href="/projets/detail.html?id=2">Caisse</a>).</strong>
        <br>J'ai utilisé le toolkit UI multiplateforme Fyne pour construire l'intégralité de l'interface utilisateur de l'application Caisse. 
        Le résultat est une application de point de vente robuste, moderne et compatible Windows, macOS et Linux, qui remplace efficacement l'ancienne version Rebol obsolète. 
        <br>Ce projet m'a permis de mettre en œuvre une technologie innovante et de la maîtriser pour livrer une solution performante adaptée aux besoins spécifiques.
    </p>

    <p>
        <strong>Conception et implémentation d'APIs RESTful performantes (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>, <a class="link" href="/projets/detail.html?id=2">Caisse</a>).</strong>
        <br>J'ai conçu et développé des APIs RESTful robustes pour servir de backend aux applications Escarcelle et Caisse. 
        Ces APIs gèrent des logiques métier complexes liées à la distribution (stocks, promotions, tarifs) et assurent une communication fluide et sécurisée entre les différentes parties du système. 
        <br>Cette expérience a renforcé ma capacité à concevoir une infrastructure fiable et scalable.
    </p>

    <p>
        <strong>Intégration de périphériques et tests unitaires (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>, <a class="link" href="/projets/detail.html?id=2">Caisse</a>).</strong>
        <br>J'ai développé des modules d'impression complets pour la génération de documents A4 et de tickets de caisse sur imprimantes thermiques (protocole ESC/POS). 
        Pour garantir la qualité et la fiabilité du code, j'ai mis en place des stratégies de tests complètes, incluant des tests unitaires et end-to-end. 
        <br>Ce travail a permis de créer un système d'impression robuste et une intégration matérielle fiable, tout en maintenant un haut niveau de qualité logicielle.
    </p>
</section>

<section class="comp-section">
    <h3><b>Autocritique</b></h3>
    <p>
        Je considère ma maîtrise de Go comme <strong>avancée</strong>. J'ai une solide compréhension de ses idiomes, de sa gestion de la concurrence (goroutines, channels) et de son écosystème. 
        C'est aujourd'hui ma compétence technique la plus forte pour le développement backend et les outils CLI. 
        J'ai pu l'approfondir rapidement grâce à des projets concrets qui m'ont permis de comprendre ses subtilités et ses bonnes pratiques.
    </p>
    <p>
        Pour les autres développeurs, je conseillerais de ne pas hésiter à explorer la documentation officielle et à bien assimiler les principes de la concurrence pour écrire du code performant et fiable. 
        Le zen de Go réside dans sa simplicité : plutôt que de transposer des patterns d'autres langages, il vaut mieux adopter sa philosophie propre.
    </p>
</section>

<section class="comp-section">
    <h3><b>Évolution</b></h3>
    <p>
        À moyen terme, je souhaite approfondir ma connaissance des frameworks web Go (comme Gin ou Echo) pour des APIs encore plus optimisées, ainsi que les déploiements sur Kubernetes. 
        Je suis également intéressé par l'intégration de Go avec des bases de données NoSQL et l'exploration de solutions d'observabilité natives (Prometheus, Grafana) pour des architectures distribuées. 
        Mon objectif est de devenir un expert capable de concevoir des systèmes Go à grande échelle.
    </p>
    <p>
        Je suis actuellement des tutoriels avancés sur l'optimisation des performances en Go et je prévois de contribuer à des projets open-source Go pour consolider mes connaissances.
    </p>
</section>

<section class="comp-section">
    <h5>Projets Associés</h5>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=2">Caisse</a></li>
        </ul>
    </div>
</section>
`
