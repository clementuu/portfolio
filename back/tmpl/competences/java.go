package tmpl

// JavaCompetence contient la description HTML de la compétence Java.
const JavaCompetence = `
<section class="comp-section">
    <h3><b>Définition</b></h3>
    <p>
        Java est un langage de programmation orienté objet, de haut niveau et robuste. 
        Sa portabilité, grâce à la machine virtuelle Java (JVM), lui permet de fonctionner sur de multiples plateformes sans recompilation. 
        Il est largement utilisé pour les applications d'entreprise, les applications Android et les systèmes distribués.
    </p>
    <p>
        J'ai utilisé Java dans le cadre de projets universitaires et personnels. C'est le tout premier langage compilé auquel j'ai été confronté, avant que je ne tombe réellement dans l'informatique. 
        Ça a pu être frustrant au début et ce n'est pas l'amour fou, mais sans lui je ne me serais peut-être moins intéressé à la programmation. 
        C'est réellement avec <a class="link" href=/competences/detail.html?id=8>Spring</a> que j'ai commencé à comprendre avec l'univers Java. 
        Dans le paysage actuel du développement, Java reste une valeur sûre pour les systèmes d'information critiques et à grande échelle.
    </p>
</section>

<section class="comp-section">
    <h3><b>Éléments de preuve</b></h3>
    <p>
        Ma compétence en Java a été principalement mise en œuvre dans le cadre du développement du Project Management Tool (PMT), où j'ai construit le backend de l'application en utilisant l'écosystème Spring Boot.
    </p>
    <p>
        <strong>Développement d'une API RESTful robuste avec Spring Boot (<a class="link" href="/projets/detail.html?id=6">PMT</a>).</strong>
        <br>Pour le projet PMT, j'ai développé une API REST complète et robuste en Java avec Spring Boot, pour gérer toutes les ressources de l'application (projets, tâches, utilisateurs). 
        Cette API sert de point d'entrée pour l'application front-end Angular, lui fournissant les données nécessaires de manière structurée et sécurisée. 
        <br>J'ai construis un backend fiable et performant, capable de gérer les opérations CRUD pour la gestion de projet. 
        <br>Ma valeur ajoutée a été de fournir une fondation solide et extensible pour l'application, garantissant une communication efficace avec le frontend.
    </p>
    <p>
        <strong>Architecture logicielle en couches et bonnes pratiques (<a class="link" href="/projets/detail.html?id=6">PMT</a>).</strong>
        <br>Le backend de PMT a été conçu selon une architecture en couches classique (Controller, Service, Repository/Store). 
        Cette approche garantit une séparation claire des responsabilités, rendant le code plus facile à maintenir, à tester et à faire évoluer. 
        <br>Le résultat est une base de code propre et modulaire, favorisant la collaboration et la scalabilité. 
        <br>Ma valeur ajoutée a été d'appliquer les principes d'ingénierie logicielle pour construire un système robuste et facile à comprendre.
    </p>
    <p>
        <strong>Intégration dans un environnement Full-Stack conteneurisé (<a class="link" href="/projets/detail.html?id=6">PMT</a>).</strong>
        <br>Le service Java/Spring de PMT s'intégrait parfaitement dans un écosystème moderne : il dialoguait avec une base de données SQL pour la persistance des données, servait de socle pour le front-end 
        <a class="link" href=/competences/detail.html?id=5>Angular</a>, et l'ensemble était orchestré et déployé dans un environnement conteneurisé avec <a class="link" href=/competences/detail.html?id=12>Docker</a>. 
        <br>Le résultat est une application full-stack complète, fonctionnelle et prête pour le déploiement. 
        <br>Ma valeur ajoutée a été de garantir une intégration fluide et une cohésion technique entre toutes les couches de l'application.
    </p>
</section>

<section class="comp-section">
    <h3><b>Autocritique</b></h3>
    <p>
        Ma maîtrise de Java est intermédiaire. J'ai une bonne compréhension des concepts orientés objet et des concepts de base de Java. 
        Mon expérience avec Spring Boot m'a permis de développer des applications backend robustes. Java est un langage que j'ai pratiqué principalement dans un cadre universitaire 
        et sur le projet PMT, ce qui m'a donné une base solide.
    </p>
    <p>
        Pour ceux qui débutent avec Java, je conseillerais de bien maîtriser les fondamentaux du langage et de la POO avant de se plonger dans les frameworks. 
        Comprendre le fonctionnement de la JVM et des threads est également un atout majeur pour écrire des applications performantes.
    </p>
</section>

<section class="comp-section">
    <h3><b>Évolution</b></h3>
    <p>
        À moyen terme, je souhaite approfondir ma connaissance des écosystèmes autour de Java, notamment les microservices avec Spring Cloud et les architectures réactives. 
        Je prévois également d'explorer les performances de Java pour des applications à haute concurrence, et son intégration avec des bases de données NoSQL. 
        Mon objectif est d'être capable de concevoir et de développer des systèmes Java distribués complexes.
    </p>
</section>

<section class="comp-section">
    <h5>Projets Associés</h5>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
