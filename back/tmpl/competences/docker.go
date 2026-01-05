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
        Cette compétence a été utilisée pour assurer la portabilité et le déploiement d'un projet majeur, 
        et fait partie de mes pratiques courantes en développement.
    </p>
    <div class="project-list">
        <strong>Projets Associés</strong>
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=2">Caisse Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
