package tmpl

// GoComp contient la description HTML de la compétence Go.
const GoComp = `
<section class="comp-section">
    <h3>Définition</h3>
    <p>
        Go, souvent appelé Golang, est un langage de programmation compilé et concurrentiel développé par Google. 
        Il est reconnu pour sa simplicité, sa performance et sa robustesse, ce qui en fait un choix privilégié 
        pour le développement de services backend, d'outils en ligne de commande et d'infrastructures cloud-native.
    </p>
    <p>
        Mon expérience avec Go couvre l'ensemble du cycle de développement d'applications, de la conception 
        d'API RESTful à la création d'applications de bureau complexes. J'apprécie particulièrement sa 
        gestion native de la concurrence, son typage statique qui garantit la sécurité du code, et son 
        écosystème d'outils très complet. Son approche minimaliste et sa capacité à produire des binaires autonomes en font un langage très pertinent dans l'écosystème des microservices et du cloud.
    </p>
</section>

<section class="comp-section">
    <h3>Éléments de preuve</h3>
    <p>
        Dans le cadre des projets <strong>Escarcelle</strong> et <strong>Caisse</strong>, 
        j'ai développé l'intégralité du backend et de l'application de caisse en Go. 
        Cette expérience a prouvé ma maîtrise de Go pour des applications critiques.
    </p>
    
    <p>
        <strong>Développement d'une application de bureau multiplateforme.</strong>
        J'ai utilisé le toolkit UI multiplateforme Fyne pour construire l'intégralité de l'interface utilisateur de l'application de caisse. 
        Le résultat a été une application de point de vente robuste, moderne et compatible Windows, macOS et Linux, résolvant le problème de l'obsolescence et de la limitation de l'ancienne version Rebol. 
        Ma valeur ajoutée a été de choisir une technologie innovante et de la maîtriser pour livrer une solution performante répondant aux besoins spécifiques du client.
    </p>

    <p>
        <strong>Conception et implémentation d'APIs RESTful performantes.</strong>
        J'ai conçu et développé des APIs RESTful robustes pour servir de backend aux applications Escarcelle et Caisse. 
        Ces APIs gèrent des logiques métier complexes liées à la distribution (stocks, promotions, tarifs). 
        Le résultat est une communication fluide et sécurisée entre les différentes parties du système, améliorant l'efficacité opérationnelle. 
        Ma valeur ajoutée a été de garantir la fiabilité et la scalabilité de l'infrastructure de données.
    </p>

    <p>
        <strong>Intégration de périphériques et tests unitaires.</strong>
        J'ai développé des modules d'impression complets pour la génération de documents A4 et de tickets de caisse sur imprimantes thermiques (protocole ESC/POS). 
        Pour garantir la qualité et la fiabilité du code, j'ai mis en place des stratégies de tests complètes, incluant des tests unitaires et end-to-end. 
        Le résultat est un système d'impression fiable et des applications robustes. Ma valeur ajoutée a été d'assurer une intégration matérielle complexe avec une qualité logicielle élevée.
    </p>
</section>

<section class="comp-section">
    <h3>Autocritique</h3>
    <p>
        Je considère ma maîtrise de Go comme <strong>avancée</strong>. J'ai une solide compréhension de ses idiomes, de sa gestion de la concurrence (goroutines, channels) et de son écosystème. 
        C'est aujourd'hui ma compétence technique la plus forte pour le développement backend et les outils CLI. 
        J'ai eu la chance de l'acquérir rapidement grâce à des projets concrets qui m'ont forcé à approfondir ses rouages, un apprentissage particulièrement stimulant et efficace.
    </p>
    <p>
        Pour les autres développeurs, je conseillerais de ne pas hésiter à plonger dans la documentation officielle qui est excellente et à bien comprendre les principes de la concurrence pour écrire du code Go performant et sans bugs. 
        Le zen de Go réside dans sa simplicité, ne cherchez pas à transposer des patterns d'autres langages, mais embrassez sa philosophie.
    </p>
</section>

<section class="comp-section">
    <h3>Évolution</h3>
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
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=2">Caisse</a></li>
        </ul>
    </div>
</section>
`
