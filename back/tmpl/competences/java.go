package tmpl

// JavaComp contient la description HTML de la compétence Java.
const JavaComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        Java est un langage de programmation orienté objet, de haut niveau et robuste. 
        Sa portabilité, grâce à la machine virtuelle Java (JVM), lui permet de fonctionner sur de multiples plateformes sans recompilation. 
        Il est largement utilisé pour les applications d'entreprise, les applications Android et les systèmes distribués.
    </p>
    <p>
        J'ai utilisé Java dans le cadre de projets universitaires et personnels, en me concentrant sur le développement backend. 
        J'ai une bonne compréhension des concepts de base de Java, de la programmation orientée objet, 
        et des librairies standards.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Ma compétence en Java a été principalement mise en œuvre dans le cadre du développement du <strong>Project Management Tool (PMT)</strong>, où j'ai construit le backend de l'application en utilisant l'écosystème <strong>Spring Boot</strong>.
    </p>
    <ul>
        <li>
            <strong>Développement d'une API RESTful :</strong> J'ai développé une API REST complète et robuste pour gérer toutes les ressources de l'application (projets, tâches, utilisateurs). Cette API sert de point d'entrée pour l'application front-end Angular, lui fournissant les données nécessaires de manière structurée et sécurisée.
        </li>
        <li>
            <strong>Architecture logicielle en couches :</strong> Le backend a été conçu selon une <strong>architecture en couches</strong> classique (Controller, Service, Repository/Store). Cette approche garantit une séparation claire des responsabilités, rendant le code plus facile à maintenir, à tester et à faire évoluer.
        </li>
        <li>
            <strong>Intégration dans un environnement Full-Stack :</strong> Le service Java/Spring s'intégrait parfaitement dans un écosystème moderne : il dialoguait avec une base de données <strong>SQL</strong> pour la persistance des données, servait de socle pour le front-end <strong>Angular</strong>, et l'ensemble était orchestré et déployé dans un environnement conteneurisé avec <strong>Docker</strong>.
        </li>
    </ul>
</section>
<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
