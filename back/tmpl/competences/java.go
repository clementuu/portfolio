package tmpl

// JavaComp contient la description HTML de la compétence Java.
const JavaComp = `
<section class="comp-section">
    <h3>Définition</h3>
    <p>
        Java est un langage de programmation orienté objet, de haut niveau et robuste. 
        Sa portabilité, grâce à la machine virtuelle Java (JVM), lui permet de fonctionner sur de multiples plateformes sans recompilation. 
        Il est largement utilisé pour les applications d'entreprise, les applications Android et les systèmes distribués.
    </p>
    <p>
        J'ai utilisé Java dans le cadre de projets universitaires et personnels, en me concentrant sur le développement backend. 
        J'ai une bonne compréhension des concepts de base de Java, de la programmation orientée objet, 
        et des librairies standards. Dans le paysage actuel du développement, Java reste une valeur sûre pour les systèmes d'information critiques et à grande échelle.
    </p>
</section>

<section class="comp-section">
    <h3>Éléments de preuve</h3>
    <p>
        Ma compétence en Java a été principalement mise en œuvre dans le cadre du développement du <strong>Project Management Tool (PMT)</strong>, où j'ai construit le backend de l'application en utilisant l'écosystème <strong>Spring Boot</strong>.
    </p>
    
    <p>
        <strong>Développement d'une API RESTful robuste avec Spring Boot.</strong>
        Pour le PMT, j'ai développé une API REST complète et robuste en Java avec Spring Boot, pour gérer toutes les ressources de l'application (projets, tâches, utilisateurs). Cette API sert de point d'entrée pour l'application front-end Angular, lui fournissant les données nécessaires de manière structurée et sécurisée. Le résultat est un backend fiable et performant, capable de gérer les opérations CRUD pour la gestion de projet. Ma valeur ajoutée a été de fournir une fondation solide et extensible pour l'application, garantissant une communication efficace avec le frontend.
    </p>

    <p>
        <strong>Architecture logicielle en couches et bonnes pratiques.</strong>
        Le backend du PMT a été conçu selon une <strong>architecture en couches</strong> classique (Controller, Service, Repository/Store). Cette approche garantit une séparation claire des responsabilités, rendant le code plus facile à maintenir, à tester et à faire évoluer. Le résultat est une base de code propre et modulaire, favorisant la collaboration et la scalabilité. Ma valeur ajoutée a été d'appliquer les principes d'ingénierie logicielle pour construire un système robuste et facile à comprendre.
    </p>

    <p>
        <strong>Intégration dans un environnement Full-Stack conteneurisé.</strong>
        Le service Java/Spring du PMT s'intégrait parfaitement dans un écosystème moderne : il dialoguait avec une base de données <strong>SQL</strong> pour la persistance des données, servait de socle pour le front-end <strong>Angular</strong>, et l'ensemble était orchestré et déployé dans un environnement conteneurisé avec <strong>Docker</strong>. Le résultat est une application full-stack complète, fonctionnelle et prête pour le déploiement. Ma valeur ajoutée a été de garantir une intégration fluide et une cohésion technique entre toutes les couches de l'application.
    </p>
</section>

<section class="comp-section">
    <h3>Autocritique</h3>
    <p>
        Ma maîtrise de Java est <strong>intermédiaire à avancé</strong>. J'ai une bonne compréhension des concepts orientés objet, des collections, et de la JVM. Mon expérience avec Spring Boot m'a permis de développer des applications backend robustes. Java est un langage que j'ai pratiqué principalement dans un cadre universitaire et sur le projet PMT, ce qui m'a donné une base solide.
    </p>
    <p>
        Pour ceux qui débutent avec Java, je conseillerais de bien maîtriser les fondamentaux du langage et de la POO avant de se plonger dans les frameworks. Comprendre le fonctionnement de la JVM et des threads est également un atout majeur pour écrire des applications performantes.
    </p>
</section>

<section class="comp-section">
    <h3>Évolution</h3>
    <p>
        À moyen terme, je souhaite approfondir ma connaissance des écosystèmes autour de Java, notamment les microservices avec <strong>Spring Cloud</strong> et les architectures réactives. Je prévois également d'explorer les performances de Java pour des applications à haute concurrence, et son intégration avec des bases de données NoSQL. Mon objectif est d'être capable de concevoir et de développer des systèmes Java distribués complexes.
    </p>
    <p>
        Je suis actuellement des tutoriels sur Spring Cloud et les architectures microservices en Java.
    </p>
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
