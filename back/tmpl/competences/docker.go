package tmpl

// DockerCompetence contient la description HTML de la compétence Docker.
const DockerCompetence = `
<section class="comp-section">
    <h3><b>Définition</b></h3>
    <p>
        Docker est une plateforme de conteneurisation permettant d'empaqueter une application et ses 
        dépendances dans un conteneur isolé, garantissant un comportement identique quel que soit 
        l'environnement d'exécution. Cette approche élimine les écarts entre postes de développement, 
        environnements de test et serveurs de production.
    </p>
    <p>
        J'utilise Docker pour standardiser mes environnements et fiabiliser mes déploiements. Grâce à 
        Docker Compose, je peux orchestrer des architectures multi-conteneurs complexes en définissant 
        les services, réseaux et volumes de manière déclarative. Cette pratique est devenue essentielle 
        dans la gestion d'applications modernes, qu'il s'agisse de microservices, de plateformes 
        full-stack ou d'écosystèmes applicatifs plus vastes comme Escarcelle.
    </p>
</section>

<section class="comp-section">
    <h3><b>Éléments de preuve</b></h3>
    <p>
        Mon expérience avec Docker couvre aussi bien la conteneurisation d'applications simples que 
        l'orchestration de plateformes complètes impliquant plusieurs services interdépendants. 
        J'ai utilisé Docker pour fiabiliser mes environnements, optimiser mes pipelines CI/CD et 
        faciliter la collaboration entre développeurs.
    </p>

    <p>
        <strong>Conteneurisation d'applications Full-Stack (Portfolio, <a class="link" href="/projets/detail.html?id=6">PMT</a>).</strong>
        Pour mon Portfolio et le projet PMT j'ai conçu des images Docker 
        capables d'embarquer à la fois le frontend (<a class="link" href=/competences/detail.html?id=4>Svelte</a>/<a class="link" href=/competences/detail.html?id=5>Angular</a>) 
        et le backend (<a class="link" href=/competences/detail.html?id=>Go</a>/<a class="link" href=/competences/detail.html?id=8>Spring</a>). 
        Cette approche a permis de garantir une reproductibilité totale entre les environnements et 
        d'accélérer les cycles de développement et de déploiement.
    </p>

    <p>
        <strong>Optimisation des builds avec Docker multi-étapes. (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>, <a class="link" href="/projets/detail.html?id=6">PMT</a>)</strong>
        J'ai systématiquement recours aux <code>Dockerfile</code> multi-étapes pour réduire la taille 
        des images finales, améliorer la sécurité et accélérer les déploiements. Par exemple, pour 
        mon portfolio, une première étape compile le frontend, une seconde compile le backend Go en 
        intégrant les assets générés, et une dernière produit une image Alpine minimale.
    </p>

    <p>
        <strong>Orchestration multi-conteneurs avancée avec Docker Compose (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>).</strong>
        Sur le projet Escarcelle, j'ai orchestré une architecture complète composée de 
        plusieurs services indépendants : frontend <a class="link" href=/competences/detail.html?id=4>Svelte</a>, backend <a class="link" href=/competences/detail.html?id=>Go</a>, 
        base de données <a class="link" href=/competences/detail.html?id=7>MySQL</a>, serveur SMTP, services internes et API partenaire. 
        Contrairement à mes projets personnels où une seule image regroupait frontend et backend, chaque composant d'Escarcelle possède sa propre image Docker, 
        permettant une modularité accrue, un déploiement granulaire et une meilleure isolation.
        <br><br>
        L'ensemble est structuré via Docker Compose afin de :
        <ul>
            <li>définir des réseaux internes pour sécuriser les échanges entre services,</li>
            <li>gérer les volumes persistants (données MySQL, logs, configurations),</li>
            <li>standardiser l'environnement de développement pour toute l'équipe,</li>
            <li>faciliter les tests d'intégration en local grâce à un écosystème complet reproductible.</li>
        </ul>
        Cette orchestration a été un élément clé pour moderniser la plateforme et accompagner la transition 
        vers un système modulaire et interopérable.
    </p>
</section>

<section class="comp-section">
    <h3><b>Autocritique</b></h3>
    <p>
        Ma maîtrise de Docker est solide. Je suis à l'aise avec la création d'images 
        optimisées, la gestion de réseaux et de volumes, l'orchestration multi-services avec Docker Compose 
        et l'intégration de Docker dans des pipelines CI/CD. Escarcelle m'a permis d'aller plus loin en 
        orchestrant un véritable écosystème applicatif composé de services hétérogènes et interdépendants.
    </p>
    <p>
        Si je devais donner un conseil à ceux qui débutent : comprendre les fondamentaux (images, 
        conteneurs, volumes, réseaux) est indispensable avant de se lancer dans des orchestrateurs plus 
        avancés comme Kubernetes. Un Dockerfile propre et une architecture Compose bien pensée 
        sont les bases d'un système conteneurisé robuste.
    </p>
</section>

<section class="comp-section">
    <h3><b>Évolution</b></h3>
    <p>
        À moyen terme, je souhaite approfondir l'orchestration avancée avec Kubernetes. 
        Après avoir géré des architectures multi-conteneurs complexes via Docker Compose, la prochaine étape 
        naturelle est la gestion d'applications distribuées à grande échelle : Helm, Operators, stratégies 
        de déploiement, observabilité, scalabilité automatique.
    </p>
    <p>
        Je suis actuellement des formations sur Kubernetes et prévois de mettre en place un cluster de 
        démonstration pour expérimenter des déploiements plus avancés, notamment en m'appuyant sur les 
        concepts déjà acquis grâce à Docker et Docker Compose.
    </p>
</section>

<section class="comp-section">
    <h5>Projets Associés</h5>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=2">Caisse</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
