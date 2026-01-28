package tmpl

// DockerComp contient la description HTML de la compétence Docker.
const DockerComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        Docker est une plateforme de conteneurisation qui permet d'empaqueter une application et ses 
        dépendances dans un conteneur logiciel isolé. Cette approche garantit que l'application 
        fonctionnera de manière uniforme et prévisible, quel que soit l'environnement sur lequel 
        elle est déployée.
    </p>
    <p>
        J'utilise Docker pour standardiser mes environnements de développement, de test et de production. 
        Cela simplifie la configuration, élimine les conflits de dépendances ("ça marche sur ma machine") 
        et facilite le déploiement continu. J'ai notamment utilisé Docker Compose pour orchestrer des 
        applications multi-conteneurs, en définissant des services, des réseaux et des volumes de manière déclarative.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Mon expérience avec Docker s'étend de la conteneurisation d'applications complètes à l'optimisation des flux CI/CD, garantissant des environnements de développement et de production cohérents et reproductibles.
    </p>
    <ul>
        <li>
            <strong>Conteneurisation d'applications Full-Stack :</strong> Pour des projets comme mon <strong>Portfolio</strong> ou <strong>PMT</strong>, j'ai mis en œuvre des architectures conteneurisées encapsulant à la fois le frontend (Svelte, Angular) et le backend (Go, Spring Boot). Cela inclut la gestion des dépendances spécifiques à chaque couche et leur compilation au sein d'une image Docker finale.
        </li>
        <li>
            <strong>Builds multi-étapes optimisés :</strong> J'utilise des <code>Dockerfile</code> multi-étapes pour réduire la taille des images finales et améliorer la sécurité. Par exemple, pour mon portfolio, une première étape compile le frontend avec Node.js, puis la seconde compile le backend Go, intégrant les assets frontaux générés, avant de produire une image finale minimale basée sur Alpine.
        </li>
        <li>
            <strong>Orchestration avec Docker Compose :</strong> Pour le développement local et les environnements de test, <strong>Docker Compose</strong> est un outil essentiel. Il me permet de définir et de lancer des applications multi-conteneurs (base de données, backend, frontend, services tiers) de manière déclarative, simplifiant ainsi la configuration et la mise en place de projets complexes.
        </li>
        <li>
            <strong>Intégration CI/CD :</strong> Docker est un pilier de mes pipelines d'intégration continue. Les conteneurs offrent un environnement de build isolé et stable, où les tests s'exécutent de manière fiable. Cela a été essentiel pour des projets comme <strong>Escarcelle</strong> et <strong>Caisse</strong>, où la portabilité de l'environnement est cruciale.
        </li>
    </ul>
    <p>
        Cette approche me permet de déployer des applications robustes et scalables, en m'assurant que l'environnement de production est une réplique fidèle de l'environnement de développement.
    </p>
</section>
<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=2">Caisse</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
